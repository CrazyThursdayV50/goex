package api

import (
	"context"

	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/request/resty"
	"github.com/CrazyThursdayV50/pkgo/trace"
)

type API[Req any, Res any] struct {
	client *resty.Client
}

func New[Req any, Res any](
	cfg *resty.Config,
	logger log.Logger,
	tracer trace.Tracer,
) *API[Req, Res] {
	opts := []resty.Option{
		resty.WithConfig(cfg),
		resty.WithLogger(logger),
	}

	if tracer != nil {
		opts = append(opts, resty.WithTracer(tracer))
	}

	client := resty.New(opts...)
	return &API[Req, Res]{
		client: client,
	}
}

func (api *API[Req, Res]) Request(ctx context.Context) {
	api.client.Request(ctx)
}
