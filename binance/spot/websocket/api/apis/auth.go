package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/auth"
)

func (api *API) Auth(ctx context.Context, data *auth.ParamsData) (result *auth.ResultData, err error) {
	var params map[string]any
	params, err = api.Sign(data)
	if err != nil {
		return
	}

	req := auth.NewRequest()
	req.Params = params

	err = request(ctx, api.api, req, &result)
	return
}
