package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/klines"
)

const (
	klinesPath = "/fapi/v1/klines"
)

func (api *API) Klines(ctx context.Context, params *klines.Params) (result klines.Result, err error) {
	err = api.none().get(ctx, klinesPath, params, &result)
	return
}
