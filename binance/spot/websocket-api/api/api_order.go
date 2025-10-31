package api

import (
	"context"

	orderModels "github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models/order"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models/order/oco"
)

type order struct {
	api *API
}

func (api *API) Order() order { return order{api} }

func (o order) Place(ctx context.Context, req *orderModels.ParamsData) (result *orderModels.ResultData, err error) {
	params := orderModels.NewParams()
	params.Params = req
	o.api.Sign(req)
	err = request(ctx, o.api, params, &result)
	return
}

func (o order) Test(ctx context.Context, req *orderModels.ParamsData) (*orderModels.ResultData, error) {
	params := orderModels.NewParamsTest()
	params.Params = req
	o.api.Sign(req)
	var result *orderModels.ResultData
	err := request(ctx, o.api, params, &result)
	return result, err
}

func (o order) PlaceMarketOrder(ctx context.Context, symbol, marketQuantity string, isBuy bool) (*orderModels.ResultData, error) {
	newOrder := orderModels.ParamsData{
		Type:   orderModels.MARKET,
		Symbol: symbol,
	}

	switch isBuy {
	case true:
		newOrder.Side = orderModels.BUY
		newOrder.QuoteOrderQty = &marketQuantity

	default:
		newOrder.Side = orderModels.SELL
		newOrder.Quantity = &marketQuantity
	}

	return o.api.Order().Place(ctx, &newOrder)
}

func (api *API) OrderList() orderList { return orderList{api} }

type orderList struct {
	api *API
}

func (o orderList) PlaceOCO(ctx context.Context, req *oco.ParamsData) (result *oco.ResultData, err error) {
	params := oco.NewParams()
	params.Params = req
	o.api.Sign(req)
	err = request(ctx, o.api, params, &result)
	return
}
