package api

import (
	"context"

	exchangeinfo "github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/exchangeInfo"
)

const (
	exchangeInfoPath = "/fapi/v1/exchangeInfo"
)

func (api *API) ExchangeInfo(ctx context.Context) (result *exchangeinfo.Result, err error) {
	err = api.none().get(ctx, exchangeInfoPath, nil, &result)
	return
}
