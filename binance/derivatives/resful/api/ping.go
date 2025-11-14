package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/ping"
)

const (
	pingPath = "/fapi/v1/ping"
)

func (api *API) Ping(ctx context.Context) (result *ping.Result, err error) {
	err = request(ctx, api, GET, pingPath, nil, &result)
	return
}
