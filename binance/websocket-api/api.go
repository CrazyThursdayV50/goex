package websocketapi

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/CrazyThursdayV50/goex/binance"
	"github.com/CrazyThursdayV50/goex/binance/models"
	"github.com/CrazyThursdayV50/goex/binance/signer"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/pkgo/builtin"
	gmap "github.com/CrazyThursdayV50/pkgo/builtin/map"
	"github.com/CrazyThursdayV50/pkgo/cron"
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type API struct {
	logger        log.Logger
	client        *binance.Client
	resultMap     builtin.MapAPI[string, *models.WsAPIResult]
	resultTimeout time.Duration
	apiKey        string
	secretKey     ed25519.PrivateKey
}

func (api *API) getResult(ctx context.Context, id string) *models.WsAPIResult {
	ctx, cancel := context.WithTimeout(ctx, api.resultTimeout)
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

func (api *API) reqId() string {
	return uuid.New().String()
}

func (api *API) Stop() { api.client.Stop() }

func (api *API) Ping(ctx context.Context) error {
	params := models.NewWsPingParams()
	params.Id = api.reqId()

	data, err := params.BinaryMarshal()
	if err != nil {
		return err
	}

	err = api.client.Send(data)
	if err != nil {
		return err
	}

	result := api.getResult(ctx, params.Id)
	if result.Status == 200 {
		return nil
	}

	return errors.New("connect failed")
}

func (api *API) ExchangeInfo(ctx context.Context, req *models.WsExchangeInfoParamsData) (*models.WsExchangeInfoResultData, error) {
	params := models.NewWsExchangeInfoParams()
	params.Id = api.reqId()
	params.Params = req

	data, err := params.BinaryMarshal()
	if err != nil {
		return nil, err
	}

	err = api.client.Send(data)
	if err != nil {
		return nil, err
	}

	result := api.getResult(ctx, params.Id)
	if result.Status == 200 {
		var data models.WsExchangeInfoResultData
		err = json.JSON().Unmarshal(result.Result, &data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}

	return nil, err
}

func (api *API) Order(ctx context.Context, req *models.WsOrderParamsData) (*models.WsOrderResultData, error) {
	params := models.NewWsOrderParams()
	params.Id = api.reqId()
	params.Params = req
	api.Sign(req)

	data, err := params.BinaryMarshal()
	if err != nil {
		return nil, err
	}

	api.logger.Infof("new order: %s", data)
	err = api.client.Send(data)
	if err != nil {
		return nil, err
	}

	result := api.getResult(ctx, params.Id)
	if result == nil {
		return nil, errors.New("request timeout")
	}

	if result.Status == 200 {
		var data models.WsOrderResultData
		err = json.JSON().Unmarshal(result.Result, &data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}

	return nil, fmt.Errorf("request failed with status: %d", result.Status)
}

func (api *API) Auth(ctx context.Context, req *models.WsAPIAuthParamsData) (*models.WsAPIAuthResultData, error) {
	params := models.NewWsAPIAuthParams()
	params.Id = api.reqId()
	params.Params = req
	api.Sign(req)

	data, err := params.BinaryMarshal()
	if err != nil {
		return nil, err
	}

	api.logger.Infof("Auth paload: %s", data)

	err = api.client.Send(data)
	if err != nil {
		return nil, err
	}

	result := api.getResult(ctx, params.Id)
	if result == nil {
		return nil, errors.New("request timeout")
	}

	if result.Status == 200 {
		var data models.WsAPIAuthResultData
		err = json.JSON().Unmarshal(result.Result, &data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}

	return nil, fmt.Errorf("request failed with status: %d", result.Status)
}

// TestOrder 发送测试下单请求
func (api *API) TestOrder(ctx context.Context, req *models.WsOrderParamsData) error {
	params := models.NewWsTestOrderParams()
	params.Id = api.reqId()
	params.Params = req
	api.Sign(req)

	data, err := params.BinaryMarshal()
	if err != nil {
		return err
	}

	api.logger.Infof("order payload: %s", data)
	err = api.client.Send(data)
	if err != nil {
		return err
	}

	result := api.getResult(ctx, params.Id)
	if result == nil {
		return errors.New("request timeout")
	}

	if result.Status != 200 {
		return fmt.Errorf("request failed with status: %d, error: %s", result.Status, result.Result)
	}

	return nil
}

// New 创建一个新的 WebSocket API 客户端
func New(ctx context.Context, logger log.Logger, apiKey, secretKey string) *API {
	resultMap := gmap.Make[string, *models.WsAPIResult](0)
	c := binance.NewClient(ctx, logger, variables.WsAPIURL(), handler(resultMap))

	c.WsClient.UpdateOptions(client.WithPingLoop(func(done <-chan struct{}, conn *websocket.Conn) {
		ctx, cancel := context.WithCancel(context.TODO())
		goo.Go(func() {
			<-done
			cancel()
		})

		cron.New(
			cron.WithContext(ctx),
			cron.WithJob(func() {
				logger.Debugf("PING sent")
				conn.WriteControl(client.PingMessage, nil, time.Now().Add(variables.WriteControlTimeout()))
			}, time.Second*5),
			cron.WithLogger(logger),
			cron.WithRunAfterStart(time.Second),
			cron.WithWaitAfterRun(false),
		).Run()
	}))

	c.Run()

	api := &API{
		client:        c,
		resultMap:     resultMap,
		logger:        logger,
		resultTimeout: variables.WsAPIReadmessageTimeout(),
	}

	// 如果提供了 API Key 和 Secret Key，自动进行身份验证
	if apiKey != "" && secretKey != "" {
		prv, err := base64.RawStdEncoding.DecodeString(secretKey)
		if err != nil {
			panic(err)
		}

		privatekey, err := x509.ParsePKCS8PrivateKey(prv)
		if err != nil {
			panic(err)
		}

		api.apiKey = apiKey
		api.secretKey = privatekey.(ed25519.PrivateKey)

		timestamp := time.Now().UnixMilli()

		var data models.WsAPIAuthParamsData
		data.ApiKey = apiKey
		data.Timestamp = timestamp

		// 发送认证请求
		authResult, err := api.Auth(ctx, &data)
		if err != nil {
			api.logger.Errorf("WebSocket API 认证失败: %v", err)
			return nil
		}

		if authResult == nil {
			api.logger.Errorf("WebSocket API 认证失败")
			return nil
		}
	}

	return api
}

func (api *API) Sign(signerData signer.SignerData) {
	signerData.SetApiKey(api.apiKey)
	signerData.SetTimestamp()
	paramsmap := signerData.Map()
	signature := signer.SignEd25519(paramsmap, api.secretKey)
	signerData.SetSignature(signature)
}
