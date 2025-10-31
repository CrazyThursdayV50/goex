package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models/account"
)

func (api *API) AccountStatus(ctx context.Context, req *account.ParamsData) (result *account.ResultData, err error) {
	params := account.NewParams()
	params.Params = req
	api.Sign(req)
	err = request(ctx, api, params, &result)
	return
}
