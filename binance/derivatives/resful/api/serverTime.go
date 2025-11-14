package api

import (
	"context"

	servertime "github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/serverTime"
)

const (
	serverTimePath = "/fapi/v1/time"
)

func (api *API) ServerTime(ctx context.Context) (result *servertime.Result, err error) {
	err = api.none().get(ctx, serverTimePath, nil, &result)
	return
}
