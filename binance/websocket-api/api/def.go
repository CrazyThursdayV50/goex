package api

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/client"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models"
	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models/auth"
	"github.com/CrazyThursdayV50/pkgo/builtin"
	gmap "github.com/CrazyThursdayV50/pkgo/builtin/map"
	"github.com/CrazyThursdayV50/pkgo/cron"
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/gorilla/websocket"
)

type API struct {
	logger          log.Logger
	client          *client.Client
	resultMap       builtin.MapAPI[string, *models.WsAPIResult]
	resultTimeout   time.Duration
	apiKey          string
	secretKeyString string

	secretKey ed25519.PrivateKey
}

func (api *API) Run(ctx context.Context) error {
	err := api.client.Run(ctx)
	if err != nil {
		return err
	}

	// 如果提供了 API Key 和 Secret Key，自动进行身份验证
	if api.apiKey != "" && api.secretKeyString != "" {
		prv, err := base64.RawStdEncoding.DecodeString(api.secretKeyString)
		if err != nil {
			return err
		}

		privatekey, err := x509.ParsePKCS8PrivateKey(prv)
		if err != nil {
			return err
		}

		api.secretKey = privatekey.(ed25519.PrivateKey)
		timestamp := time.Now().UnixMilli()

		var data auth.ParamsData
		data.ApiKey = api.apiKey
		data.Timestamp = timestamp

		// 发送认证请求
		authResult, err := api.Auth(ctx, &data)
		if err != nil {
			api.logger.Errorf("WebSocket API 认证失败: %v", err)
			return err
		}

		if authResult == nil {
			api.logger.Errorf("WebSocket API 认证失败")
			return err
		}
	}

	return nil
}

// New 创建一个新的 WebSocket API 客户端
func New(logger log.Logger, apiKey, secretKey string) *API {
	resultMap := gmap.Make[string, *models.WsAPIResult](0)

	pingLoop := func(done <-chan struct{}, conn *websocket.Conn) {
		ctx, cancel := context.WithCancel(context.TODO())
		goo.Go(func() {
			<-done
			cancel()
		})

		cron.New(
			cron.WithJob(func() {
				logger.Debugf("PING sent")
				conn.WriteControl(PingMessage, nil, time.Now().Add(variables.WriteControlTimeout()))
			}, time.Second*10),
			cron.WithLogger(logger),
			cron.WithRunAfterStart(time.Second),
			cron.WithWaitAfterRun(false),
		).Run(ctx)
	}

	c := client.NewClient(logger, variables.WsAPIURL(), handler(resultMap), pingLoop)

	api := &API{
		client:          c,
		resultMap:       resultMap,
		logger:          logger,
		resultTimeout:   variables.WsAPIReadmessageTimeout(),
		apiKey:          apiKey,
		secretKeyString: secretKey,
	}

	return api
}

func (api *API) Stop() { api.client.Stop() }
