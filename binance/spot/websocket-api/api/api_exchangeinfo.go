package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models/exchangeinfo"
)

func (api *API) ExchangeInfo(ctx context.Context, req *exchangeinfo.ParamsData) (result *exchangeinfo.ResultData, err error) {
	params := exchangeinfo.NewParams()
	params.Params = req
	err = request(ctx, api, params, &result)
	return
}
