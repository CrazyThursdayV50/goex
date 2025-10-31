package apis

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"
	"github.com/CrazyThursdayV50/goex/binance/iface"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/goex/infra/sign"
	"github.com/CrazyThursdayV50/goex/infra/wsclient"
	"github.com/CrazyThursdayV50/pkgo/builtin"
	gmap "github.com/CrazyThursdayV50/pkgo/builtin/map"
	"github.com/CrazyThursdayV50/pkgo/cron"
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
	"github.com/CrazyThursdayV50/pkgo/worker"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type API struct {
	logger   log.Logger
	wsclient *wsclient.Client

	resultChan chan *models.Result
	resultMap  builtin.MapAPI[string, *models.Result]

	apikey       string
	secretString string

	secretKey ed25519.PrivateKey
}

func (api *API) reqID() string {
	return uuid.New().String()
}

func (api *API) Sign(signerData iface.SignerData) {
	signerData.SetAPIKEY(api.apikey)
	signerData.SetTimestamp()
	paramsmap := signerData.Map()
	signature := sign.Ed25519(paramsmap, api.secretKey)
	signerData.SetSignature(signature)
}

func (api *API) GetResult(ctx context.Context, id string) *models.Result {
	ctx, cancel := context.WithTimeout(ctx, derivatives.WsAPI().ReadMessageTimeout())
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			if result := api.resultMap.Get(id).Unwrap(); result != nil {
				return result
			}

			time.Sleep(time.Millisecond)
		}
	}
}

var PingMessage = client.PingMessage

// New 创建一个新的 WebSocket API 客户端
func New(logger log.Logger, apiKey, secretKey string) *API {
	resultMap := gmap.Make[string, *models.Result](0)

	pingLoop := func(done <-chan struct{}, conn *websocket.Conn) {
		ctx, cancel := context.WithCancel(context.TODO())
		goo.Go(func() {
			<-done
			cancel()
		})

		cron.New(
			cron.WithJob(func() {
				logger.Debugf("Send: PING")
				conn.WriteControl(PingMessage, nil, time.Now().Add(variables.WriteControlTimeout()))
			}, time.Second*30),
			cron.WithLogger(logger),
			cron.WithRunAfterStart(time.Second),
			cron.WithWaitAfterRun(false),
		).Run(ctx)
	}

	ch := make(chan *models.Result, 100)
	c := wsclient.NewClient(
		logger,
		derivatives.WsAPI().Endpoint(), variables.GetProxy(),
		handler(ch),
		pingLoop,
	)

	api := &API{
		wsclient:     c,
		resultChan:   ch,
		resultMap:    resultMap,
		logger:       logger,
		apikey:       apiKey,
		secretString: secretKey,
	}

	return api
}

func handler(ch chan<- *models.Result) func(context.Context, log.Logger, int, []byte, func(error)) (int, []byte) {
	return func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
		var result models.Result
		err := json.JSON().Unmarshal(b, &result)
		if err != nil {
			l.Errorf("unmarshal failed: %v", err)
			return client.TextMessage, nil
		}

		select {
		case <-ctx.Done():
		case ch <- &result:
		}

		return client.TextMessage, nil
	}
}

func (api *API) Run(ctx context.Context) error {
	err := api.wsclient.Run(ctx)
	if err != nil {
		return err
	}

	worker, _ := worker.New("WsResultReceive", func(r *models.Result) {
		api.resultMap.AddSoft(r.ID, r)
	})
	worker.WithLogger(api.logger)
	worker.WithGraceful(true)
	worker.WithTrigger(api.resultChan)
	worker.Run(ctx)

	// 如果提供了 API Key 和 Secret Key，自动进行身份验证
	if api.apikey != "" && api.secretString != "" {
		prv, err := base64.RawStdEncoding.DecodeString(api.secretString)
		if err != nil {
			return err
		}

		privatekey, err := x509.ParsePKCS8PrivateKey(prv)
		if err != nil {
			return err
		}

		api.secretKey = privatekey.(ed25519.PrivateKey)

		// 发送认证请求
		logonData, err := api.Session().Logon(ctx)
		if err != nil {
			api.logger.Errorf("WebSocket API 认证失败: %v", err)
			return err
		}

		if logonData == nil {
			api.logger.Errorf("WebSocket API 认证失败")
			return err
		}

		api.logger.Infof("logon success: %+v", logonData)
	}

	return nil
}

func request[reqData any, resultData any](ctx context.Context, api *API, req *models.Request[reqData], dst *resultData) error {
	req.ID = api.reqID()

	data, err := json.JSON().Marshal(req)
	if err != nil {
		return err
	}

	err = api.wsclient.Send(data)
	if err != nil {
		return err
	}

	result := api.GetResult(ctx, req.ID)
	if result == nil {
		return errors.New("request timeout")
	}

	if result.Status == 200 {
		err = json.JSON().Unmarshal(result.Result, dst)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("request failed with status: %d, error: %s", result.Status, result.Error)
}
