package api

import (
	"context"
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models"
	"github.com/CrazyThursdayV50/goex/binance/sign"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/goex/infra/utils"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/request/resty"
	"github.com/CrazyThursdayV50/pkgo/trace"
)

/*
 NONE	不需要鉴权的接口
 TRADE	需要有效的API-KEY和签名
 USER_DATA	需要有效的API-KEY和签名
 USER_STREAM	需要有效的API-KEY
 MARKET_DATA	需要有效的API-KEY
*/

type API struct {
	baseURL string
	client  *resty.Client

	apiKey     string
	signerFunc func(any) (string, error)
}

func New(
	cfg *resty.Config,
	logger log.Logger,
	tracer trace.Tracer,
	apikey, secret string,
) (*API, error) {
	opts := []resty.Option{
		resty.WithConfig(cfg),
		resty.WithLogger(logger),
	}

	if tracer != nil {
		opts = append(opts, resty.WithTracer(tracer))
	}

	client := resty.New(opts...)

	private, err := sign.ParseSecretEd25519(apikey, secret)
	if err != nil {
		return nil, err
	}

	signerFunc := sign.NewSignerFuncEd25519QueryParams(apikey, private)

	return &API{
		baseURL:    derivatives.Rest().Endpoint(),
		client:     client,
		apiKey:     apikey,
		signerFunc: signerFunc,
	}, nil
}

const (
	GET    = "GET"
	PUT    = "PUT"
	PATCH  = "PATCH"
	POST   = "POST"
	DELETE = "DELETE"
)

const paramsKeySignature = "signature"

func get(ctx context.Context, api *API, url, queries string, result any) error {
	return request(ctx, api, GET, fmt.Sprintf("%s?%s", url, queries), nil, result)
}

func post(ctx context.Context, api *API, url string, body any, result any) error {
	return request(ctx, api, POST, url, nil, result)
}

func request(ctx context.Context, api *API, method, path string, params map[string]any, result any) error {
	var resultError *models.ResultError
	r := api.client.Request(ctx).
		SetError(&resultError).
		SetResult(&result)

	url := api.baseURL + path
	r.SetHeader("X-MBX-APIKEY", api.apiKey)

	switch method {
	case GET:
		for k, v := range params {
			r.SetQueryParam(k, fmt.Sprintf("%v", v))
		}

	default:
		r.SetBody(params)
	}

	_, err := r.Execute(method, url)
	if err != nil {
		return err
	}

	if resultError != nil {
		return resultError
	}

	return nil
}

type TRADE struct {
	baseURL    string
	apikey     string
	client     *resty.Client
	signerFunc func(any) (string, error)
}

func (api *API) trade() *TRADE {
	return &TRADE{
		baseURL:    api.baseURL,
		apikey:     api.apiKey,
		client:     api.client,
		signerFunc: api.signerFunc,
	}
}

type USERDATA = TRADE

func (api *API) userdata() *USERDATA {
	return api.trade()
}

func (n *TRADE) get(ctx context.Context, path string, params any, result any) error {
	url := n.baseURL + path

	payload, err := n.signerFunc(params)
	if err != nil {
		return err
	}

	var resultError *models.ResultError
	_, err = n.client.Request(ctx).
		SetError(&resultError).
		SetResult(result).
		SetHeader("X-MBX-APIKEY", n.apikey).
		Get(fmt.Sprintf("%s?%s", url, payload))
	if err != nil {
		return err
	}

	if resultError != nil {
		return resultError
	}

	return nil
}

func (n *TRADE) post(ctx context.Context, path string, params any, result any) error {
	url := n.baseURL + path

	payload, err := n.signerFunc(params)
	if err != nil {
		return err
	}

	var resultError *models.ResultError
	_, err = n.client.Request(ctx).
		SetError(&resultError).
		SetResult(result).
		SetHeader("X-MBX-APIKEY", n.apikey).
		SetBody(payload).
		Post(url)
	if err != nil {
		return err
	}

	if resultError != nil {
		return resultError
	}

	return nil
}

func (api *API) none() *NONE {
	return &NONE{
		baseURL: api.baseURL,
		client:  api.client,
	}
}

type NONE struct {
	baseURL string
	client  *resty.Client
}

func (n *NONE) get(ctx context.Context, path string, params any, result any) error {
	url := n.baseURL + path

	queriesMap, err := utils.MapString(params)
	if err != nil {
		return err
	}

	var resultError *models.ResultError
	_, err = n.client.Request(ctx).
		SetError(&resultError).
		SetResult(result).
		SetQueryParams(queriesMap).
		Get(url)
	if err != nil {
		return err
	}

	if resultError != nil {
		return resultError
	}

	return nil
}

type USERSTREAM struct {
	baseURL string
	apikey  string
	client  *resty.Client
}

func (api *API) userstream() *USERSTREAM {
	return &USERSTREAM{
		baseURL: api.baseURL,
		apikey:  api.apiKey,
		client:  api.client,
	}
}

func (n *USERSTREAM) post(ctx context.Context, path string, params any, result any) error {
	url := n.baseURL + path

	var resultError *models.ResultError
	_, err := n.client.Request(ctx).
		SetError(&resultError).
		SetResult(result).
		SetHeader("X-MBX-APIKEY", n.apikey).
		Post(url)
	if err != nil {
		return err
	}

	if resultError != nil {
		return resultError
	}

	return nil
}

func (n *USERSTREAM) put(ctx context.Context, path string, params any, result any) error {
	url := n.baseURL + path

	var resultError *models.ResultError
	_, err := n.client.Request(ctx).
		SetError(&resultError).
		SetResult(result).
		SetHeader("X-MBX-APIKEY", n.apikey).
		Put(url)
	if err != nil {
		return err
	}

	if resultError != nil {
		return resultError
	}

	return nil
}

func (n *USERSTREAM) delete(ctx context.Context, path string, params any, result any) error {
	url := n.baseURL + path

	var resultError *models.ResultError
	_, err := n.client.Request(ctx).
		SetError(&resultError).
		SetResult(result).
		SetHeader("X-MBX-APIKEY", n.apikey).
		Delete(url)
	if err != nil {
		return err
	}

	if resultError != nil {
		return resultError
	}

	return nil
}
