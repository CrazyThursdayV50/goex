package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/klines"
	"github.com/CrazyThursdayV50/goex/infra/utils"
)

type klineRequest struct {
	api  *API
	data *klines.ParamsData
}

func (r *klineRequest) Symbol(value string) *klineRequest {
	r.data.Symbol = value
	return r
}

func (r *klineRequest) StartTime(value uint64) *klineRequest {
	r.data.StartTime = value
	return r
}

func (r *klineRequest) EndTime(value uint64) *klineRequest {
	r.data.EndTime = value
	return r
}

func (r *klineRequest) Interval(value string) *klineRequest {
	r.data.Interval = value
	return r
}

func (r *klineRequest) TimeZone(value string) *klineRequest {
	r.data.TimeZone = value
	return r
}

func (r *klineRequest) Limit(value int) *klineRequest {
	r.data.Limit = value
	return r
}

func (r *klineRequest) Do(ctx context.Context) (result *klines.ResultData, err error) {
	req := klines.NewParams()
	req.Params, _ = utils.MapAny(r.data)
	err = request(ctx, r.api.api, req, &result)
	return
}

func (api *API) Klines() *klineRequest {
	return &klineRequest{api: api, data: new(klines.ParamsData)}
}
