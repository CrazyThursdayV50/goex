package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/klines"
	"github.com/CrazyThursdayV50/goex/infra/utils"
)

const (
	klinesPath = "/fapi/v1/klines"
)

func (api *API) Klines(ctx context.Context, params *klines.Params) (result *klines.Result, err error) {
	paramsMap, err := utils.MapString(params)
	if err != nil {
		return nil, err
	}

	err = request(ctx, api, GET, klinesPath, paramsMap, &result)
	return
}
