package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models/auth"
)

func (api *API) Auth(ctx context.Context, req *auth.ParamsData) (result *auth.ResultData, err error) {
	params := auth.NewParams()
	params.Params = req
	api.Sign(req)
	err = request(ctx, api, params, &result)
	return
}
