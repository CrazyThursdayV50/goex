package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models/klines"
)

type klineRequest struct {
	api    *API
	params *klines.Params
}

func (r *klineRequest) Symbol(value string) *klineRequest {
	r.params.Params.Symbol = value
	return r
}

func (r *klineRequest) StartTime(value uint64) *klineRequest {
	r.params.Params.StartTime = value
	return r
}

func (r *klineRequest) EndTime(value uint64) *klineRequest {
	r.params.Params.EndTime = value
	return r
}

func (r *klineRequest) Interval(value string) *klineRequest {
	r.params.Params.Interval = value
	return r
}

func (r *klineRequest) TimeZone(value string) *klineRequest {
	r.params.Params.TimeZone = value
	return r
}

func (r *klineRequest) Limit(value int) *klineRequest {
	r.params.Params.Limit = value
	return r
}

func (r *klineRequest) Do(ctx context.Context) (result klines.ResultData, err error) {
	err = request(ctx, r.api, r.params, &result)
	return
}

func (api *API) Klines() *klineRequest {
	params := klines.NewParams()
	return &klineRequest{api: api, params: params}
}
