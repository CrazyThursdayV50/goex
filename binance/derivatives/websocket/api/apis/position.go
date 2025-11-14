package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/position"
)

func (api *API) Position(ctx context.Context, data *position.RequestData) (result position.ResultData, err error) {
	req := position.Params()
	req.Params, err = api.Sign(data)
	if err != nil {
		return
	}

	err = request(ctx, api.api, req, &result)
	return
}
