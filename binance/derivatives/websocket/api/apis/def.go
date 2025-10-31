package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"
	"github.com/CrazyThursdayV50/goex/binance/sign"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/goex/infra/iface"
	"github.com/CrazyThursdayV50/goex/infra/websocket/api"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
)

type API struct {
	api *api.API[*models.Request, *models.Result]
}

func New(logger log.Logger, apikey, secret string) (*API, error) {
	var signerFunc iface.SignerFunc

	if apikey != "" && secret != "" {
		secretKey, err := sign.ParseSecretEd25519(apikey, secret)
		if err != nil {
			return nil, err
		}
		signerFunc = sign.NewSignerFuncEd25519(apikey, secretKey)
	}

	api := api.New[*models.Request, *models.Result](
		logger,
		derivatives.WsAPI().Endpoint(),
		signerFunc,
		Handler,
	)

	return &API{api}, nil
}

func Handler(ch chan *models.Result) client.MessageHandler {
	return func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
		var result models.Result
		err := json.JSON().Unmarshal(b, &result)
		if err != nil {
			l.Errorf("unmarshal failed: %v", err)
			return client.TextMessage, nil
		}

		select {
		case <-ctx.Done():
			close(ch)

		case ch <- &result:
		}

		return client.TextMessage, nil
	}
}

func (api *API) Run(ctx context.Context) error {
	return api.api.Run(ctx)
}

func (api *API) Stop() {
	api.api.Stop()
}

func (api *API) Sign(requestData any) (map[string]any, error) {
	return api.
		api.
		Sign(requestData)
}

// func (api *API) GetResult(ctx context.Context, id string) *models.Result {
// 	ctx, cancel := context.WithTimeout(ctx, derivatives.WsAPI().ReadMessageTimeout())
// 	defer cancel()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil

// 		default:
// 			if result := api.resultMap.Get(id).Unwrap(); result != nil {
// 				return result
// 			}

// 			time.Sleep(time.Millisecond)
// 		}
// 	}
// }

// var PingMessage = client.PingMessage

// // New 创建一个新的 WebSocket API 客户端
// func New(logger log.Logger, apiKey, secretKey string) *API {
// 	resultMap := gmap.Make[string, *models.Result](0)

// 	pingLoop := func(done <-chan struct{}, conn *websocket.Conn) {
// 		ctx, cancel := context.WithCancel(context.TODO())
// 		goo.Go(func() {
// 			<-done
// 			cancel()
// 		})

// 		cron.New(
// 			cron.WithJob(func() {
// 				logger.Debugf("Send: PING")
// 				conn.WriteControl(PingMessage, nil, time.Now().Add(variables.WriteControlTimeout()))
// 			}, time.Second*30),
// 			cron.WithLogger(logger),
// 			cron.WithRunAfterStart(time.Second),
// 			cron.WithWaitAfterRun(false),
// 		).Run(ctx)
// 	}

// 	ch := make(chan *models.Result, 100)
// 	c := client1.NewClient(
// 		logger,
// 		derivatives.WsAPI().Endpoint(), variables.GetProxy(),
// 		handler(ch),
// 		pingLoop,
// 	)

// 	api := &API{
// 		wsclient:     c,
// 		resultChan:   ch,
// 		resultMap:    resultMap,
// 		logger:       logger,
// 		apikey:       apiKey,
// 		secretString: secretKey,
// 	}

// 	return api
// }

// func handler(ch chan<- *models.Result) func(context.Context, log.Logger, int, []byte, func(error)) (int, []byte) {
// 	return func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
// 		var result models.Result
// 		err := json.JSON().Unmarshal(b, &result)
// 		if err != nil {
// 			l.Errorf("unmarshal failed: %v", err)
// 			return client.TextMessage, nil
// 		}

// 		select {
// 		case <-ctx.Done():
// 			close(ch)

// 		case ch <- &result:
// 		}

// 		return client.TextMessage, nil
// 	}
// }

// func (api *API) Run(ctx context.Context) error {
// 	err := api.wsclient.Run(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	worker, _ := worker.New("WsResultReceive", func(r *models.Result) {
// 		api.resultMap.AddSoft(r.ID, r)
// 	})
// 	worker.WithLogger(api.logger)
// 	worker.WithGraceful(true)
// 	worker.WithTrigger(api.resultChan)
// 	worker.Run(ctx)

// 	// 如果提供了 API Key 和 Secret Key，自动进行身份验证
// 	if api.apikey != "" && api.secretString != "" {
// 		prv, err := base64.RawStdEncoding.DecodeString(api.secretString)
// 		if err != nil {
// 			return err
// 		}

// 		privatekey, err := x509.ParsePKCS8PrivateKey(prv)
// 		if err != nil {
// 			return err
// 		}

// 		api.secretKey = privatekey.(ed25519.PrivateKey)

// 		// 发送认证请求
// 		logonData, err := api.Session().Logon(ctx)
// 		if err != nil {
// 			api.logger.Errorf("WebSocket API 认证失败: %v", err)
// 			return err
// 		}

// 		if logonData == nil {
// 			api.logger.Errorf("WebSocket API 认证失败")
// 			return err
// 		}

// 		api.logger.Infof("logon success: %+v", logonData)
// 	}

// 	return nil
// }

// func request[reqData any, resultData any](ctx context.Context, api *API, req *models.Request[reqData], dst *resultData) error {
// 	req.ID = api.reqID()

// 	data, err := json.JSON().Marshal(req)
// 	if err != nil {
// 		return err
// 	}

// 	err = api.wsclient.Send(data)
// 	if err != nil {
// 		return err
// 	}

// 	result := api.GetResult(ctx, req.ID)
// 	if result == nil {
// 		return errors.New("request timeout")
// 	}

// 	if result.Status == 200 {
// 		err = json.JSON().Unmarshal(result.Result, dst)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}

// 	return fmt.Errorf("request failed with status: %d, error: %s", result.Status, result.Err)
// }
