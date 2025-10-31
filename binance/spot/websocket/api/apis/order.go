package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/order"
)

type orderRouter struct {
	api *API
}

func (api *API) Order() orderRouter { return orderRouter{api} }

func (o orderRouter) Place(ctx context.Context, data *order.ParamsData) (result *order.ResultData, err error) {
	var params map[string]any
	params, err = o.api.Sign(data)
	if err != nil {
		return
	}

	req := order.NewParams()
	req.Params = params

	err = request(ctx, o.api.api, req, &result)
	return
}

func (o orderRouter) Test(ctx context.Context, data *order.ParamsData) (result *order.ResultData, err error) {
	var params map[string]any
	params, err = o.api.Sign(data)
	if err != nil {
		return
	}

	req := order.NewParamsTest()
	req.Params = params

	err = request(ctx, o.api.api, req, &result)
	return result, err
}

func (o orderRouter) PlaceMarketOrder(ctx context.Context, symbol, marketQuantity string, isBuy bool) (result *order.ResultData, err error) {
	newOrder := order.ParamsData{
		Type:   order.MARKET,
		Symbol: symbol,
	}

	switch isBuy {
	case true:
		newOrder.Side = order.BUY
		newOrder.QuoteOrderQty = &marketQuantity

	default:
		newOrder.Side = order.SELL
		newOrder.Quantity = &marketQuantity
	}

	return o.api.Order().Place(ctx, &newOrder)
}
