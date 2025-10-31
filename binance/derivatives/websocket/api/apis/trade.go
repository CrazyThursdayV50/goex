package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/trade"
	"github.com/CrazyThursdayV50/goex/infra/utils"
)

type Trade struct{ api *API }

func (api *API) Trade() *Trade {
	return &Trade{api: api}
}

func (t *Trade) PlaceOrder(ctx context.Context, data *trade.PlaceData) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	data.SetTimestamp()
	req.Params = data
	err = request(ctx, t.api, req, &result)
	return
}

// func (t *Trade) PlaceOrderTest(ctx context.Context, data *trade.PlaceData) (result *trade.PlaceResultData, err error) {
// 	req := trade.PlaceTest()
// 	t.api.Sign(data)
// 	req.Params = data
// 	err = request(ctx, t.api, req, &result)
// 	return
// }

type custom struct {
	api *API
}

// 自定义接口
// 为了方便实用，简化了 binance API
func (t *Trade) Custom() custom {
	return custom{api: t.api}
}

// OpenLongMarketSingle 单向持仓市价开多
func (c custom) OpenLongMarketSingle(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = &trade.PlaceData{
		Symbol:   symbol,
		Quantity: utils.Ptr(quantity),

		OrderSide:        trade.SIDE_BUY,
		OrderType:        trade.TYPE_MARKET,
		NewOrderRespType: utils.Ptr(trade.NEW_ORDER_RESP_TYPE_RESULT),
	}

	err = request(ctx, c.api, req, &result)
	return
}

// OpenShortMarketSingle 单向持仓市价开空
func (c custom) OpenShortMarketSingle(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = &trade.PlaceData{
		Symbol:   symbol,
		Quantity: utils.Ptr(quantity),

		OrderSide:        trade.SIDE_SELL,
		OrderType:        trade.TYPE_MARKET,
		NewOrderRespType: utils.Ptr(trade.NEW_ORDER_RESP_TYPE_RESULT),
	}

	err = request(ctx, c.api, req, &result)
	return
}

// ReduceShortMarketSingle 单向持仓 看空仓位 市价减仓
func (c custom) ReduceShortMarketSingle(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = &trade.PlaceData{
		Symbol:   symbol,
		Quantity: utils.Ptr(quantity),

		OrderSide:        trade.SIDE_BUY,
		OrderType:        trade.TYPE_MARKET,
		ReduceOnly:       utils.Ptr(trade.REDUCE_ONLY_TRUE),
		NewOrderRespType: utils.Ptr(trade.NEW_ORDER_RESP_TYPE_RESULT),
	}

	err = request(ctx, c.api, req, &result)
	return
}

// ReduceLongMarketSingle 单向持仓 看多仓位 市价减仓
func (c custom) ReduceLongMarketSingle(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = &trade.PlaceData{
		Symbol:   symbol,
		Quantity: utils.Ptr(quantity),

		OrderSide:        trade.SIDE_SELL,
		OrderType:        trade.TYPE_MARKET,
		ReduceOnly:       utils.Ptr(trade.REDUCE_ONLY_TRUE),
		NewOrderRespType: utils.Ptr(trade.NEW_ORDER_RESP_TYPE_RESULT),
	}

	err = request(ctx, c.api, req, &result)
	return
}
