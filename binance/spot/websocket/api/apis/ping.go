package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/ping"
)

func (api *API) Ping(ctx context.Context) (result *ping.ResultData, err error) {
	req := ping.NewParams()
	err = request(ctx, api.api, req, &result)
	return
}
