package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/order/oco"
)

func (api *API) OrderList() orderListRouter { return orderListRouter{api} }

type orderListRouter struct {
	api *API
}

func (o orderListRouter) PlaceOCO(ctx context.Context, data *oco.ParamsData) (result *oco.ResultData, err error) {
	var params map[string]any
	params, err = o.api.Sign(data)
	if err != nil {
		return
	}

	req := oco.NewParams()
	req.Params = params

	err = request(ctx, o.api.api, req, &result)
	return
}
