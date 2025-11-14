package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/request/resty"
	"github.com/CrazyThursdayV50/pkgo/trace"
)

type API struct {
	baseURL string
	client  *resty.Client
}

func New(
	cfg *resty.Config,
	logger log.Logger,
	tracer trace.Tracer,
) *API {
	opts := []resty.Option{
		resty.WithConfig(cfg),
		resty.WithLogger(logger),
	}

	if tracer != nil {
		opts = append(opts, resty.WithTracer(tracer))
	}

	client := resty.New(opts...)
	return &API{
		baseURL: derivatives.Rest().Endpoint(),
		client:  client,
	}
}

const (
	GET    = "GET"
	PUT    = "PUT"
	PATCH  = "PATCH"
	POST   = "POST"
	DELETE = "DELETE"
)

func request(ctx context.Context, api *API, method, path string, params map[string]string, result any) error {
	var resultError *models.ResultError
	r := api.client.Request(ctx).
		SetError(&resultError).
		SetResult(&result)

	url := api.baseURL + path
	switch method {
	case GET:
		r.SetQueryParams(params)

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
