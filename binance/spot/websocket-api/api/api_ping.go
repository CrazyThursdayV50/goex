package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models/ping"
)

func (api *API) Ping(ctx context.Context) (result ping.ResultData, err error) {
	params := ping.NewParams()
	params.Id = api.reqID()
	err = request(ctx, api, params, &result)
	return
}
