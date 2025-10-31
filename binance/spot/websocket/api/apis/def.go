package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/sign"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models"
	"github.com/CrazyThursdayV50/goex/binance/variables/spot"
	"github.com/CrazyThursdayV50/goex/infra/iface"
	"github.com/CrazyThursdayV50/goex/infra/websocket/api"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type API struct {
	api *api.API[*models.WsAPIRequest, *models.WsAPIResult]
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

	api := api.New[*models.WsAPIRequest, *models.WsAPIResult](
		logger,
		spot.WsAPI().Endpoint(),
		signerFunc,
		Handler,
	)

	return &API{api}, nil
}

func (api *API) Run(ctx context.Context) error {
	return api.api.Run(ctx)
}

func (api *API) Stop() {
	api.api.Stop()
}

func (api *API) Sign(requestData any) (map[string]any, error) {
	return api.api.Sign(requestData)
}

func Handler(ch chan *models.WsAPIResult) client.MessageHandler {
	return func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
		var result models.WsAPIResult
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

// type API struct {
// 	logger          log.Logger
// 	client          *client1.Client
// 	resultMap       builtin.MapAPI[string, *models.WsAPIResult]
// 	resultTimeout   time.Duration
// 	apiKey          string
// 	secretKeyString string

// 	secretKey ed25519.PrivateKey
// }

// func (api *API) Run(ctx context.Context) error {
// 	err := api.client.Run(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	// 如果提供了 API Key 和 Secret Key，自动进行身份验证
// 	if api.apiKey != "" && api.secretKeyString != "" {
// 		prv, err := base64.RawStdEncoding.DecodeString(api.secretKeyString)
// 		if err != nil {
// 			return err
// 		}

// 		privatekey, err := x509.ParsePKCS8PrivateKey(prv)
// 		if err != nil {
// 			return err
// 		}

// 		api.secretKey = privatekey.(ed25519.PrivateKey)
// 		timestamp := time.Now().UnixMilli()

// 		var data auth.ParamsData
// 		data.ApiKey = api.apiKey
// 		data.Timestamp = timestamp

// 		// 发送认证请求
// 		authResult, err := api.Auth(ctx, &data)
// 		if err != nil {
// 			api.logger.Errorf("WebSocket API 认证失败: %v", err)
// 			return err
// 		}

// 		if authResult == nil {
// 			api.logger.Errorf("WebSocket API 认证失败")
// 			return err
// 		}
// 	}

// 	return nil
// }

// var PingMessage = client.PingMessage

// // New 创建一个新的 WebSocket API 客户端
// func New(logger log.Logger, apiKey, secretKey string) *API {
// 	resultMap := gmap.Make[string, *models.WsAPIResult](0)

// 	pingLoop := func(done <-chan struct{}, conn *websocket.Conn) {
// 		ctx, cancel := context.WithCancel(context.TODO())
// 		goo.Go(func() {
// 			<-done
// 			cancel()
// 		})

// 		cron.New(
// 			cron.WithJob(func() {
// 				logger.Debugf("PING sent")
// 				conn.WriteControl(PingMessage, nil, time.Now().Add(variables.WriteControlTimeout()))
// 			}, time.Second*10),
// 			cron.WithLogger(logger),
// 			cron.WithRunAfterStart(time.Second),
// 			cron.WithWaitAfterRun(false),
// 		).Run(ctx)
// 	}

// 	c := client1.NewClient(
// 		logger,
// 		spot.WsAPI().Endpoint(), variables.GetProxy(),
// 		handler(resultMap),
// 		pingLoop,
// 	)

// 	api := &API{
// 		client:          c,
// 		resultMap:       resultMap,
// 		logger:          logger,
// 		resultTimeout:   spot.WsAPI().ReadMessageTimeout(),
// 		apiKey:          apiKey,
// 		secretKeyString: secretKey,
// 	}

// 	return api
// }

// func (api *API) Stop() { api.client.Stop() }
