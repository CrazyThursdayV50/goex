package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/account"
)

func (api *API) AccountStatus(ctx context.Context, data *account.ParamsData) (result *account.ResultData, err error) {
	var params map[string]any
	params, err = api.Sign(data)
	if err != nil {
		return
	}

	req := account.NewRequest()
	req.Params = params

	err = request(ctx, api.api, req, &result)
	return
}
