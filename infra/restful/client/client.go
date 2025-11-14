package client

import (
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/request/resty"
	"github.com/CrazyThursdayV50/pkgo/trace"
)

type Client struct {
	client *resty.Client
}

func New(
	cfg *resty.Config,
	logger log.Logger,
	tracer trace.Tracer,
) *Client {
	opts := []resty.Option{
		resty.WithConfig(cfg),
		resty.WithLogger(logger),
	}

	if tracer != nil {
		opts = append(opts, resty.WithTracer(tracer))
	}

	client := resty.New(opts...)
	return &Client{client: client}
}
