package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/trade"
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

type custom struct {
	api *API
}

// 自定义接口
// 为了方便实用，简化了 binance API
func (t *Trade) Custom() custom {
	return custom{api: t.api}
}

// OpenLongMarketSingle 单向持仓市价开多
func (c custom) SingleOpenLongMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = new(trade.PlaceData)
	req.Params.SingleOpenLongMarket(symbol, quantity)
	err = request(ctx, c.api, req, &result)
	return
}

// OpenShortMarketSingle 单向持仓市价开空
func (c custom) SingleOpenShortMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = new(trade.PlaceData)
	req.Params.SingleOpenShortMarket(symbol, quantity)
	err = request(ctx, c.api, req, &result)
	return
}

// ReduceLongMarketSingle 单向持仓 看多仓位 市价减仓
func (c custom) SingleReduceLongMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = new(trade.PlaceData)
	req.Params.SingleReduceLongMarket(symbol, quantity)

	err = request(ctx, c.api, req, &result)
	return
}

// ReduceShortMarketSingle 单向持仓 看空仓位 市价减仓
func (c custom) SingleReduceShortMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = new(trade.PlaceData)
	req.Params.SingleReduceShortMarket(symbol, quantity)

	err = request(ctx, c.api, req, &result)
	return
}

// SingleLongTakeProfitMarket 单向持仓 看多仓位 市价止盈
func (c custom) SingleLongTakeProfitMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = new(trade.PlaceData)
	req.Params.SingleLongTakeProfitMarket(symbol, quantity, stopPrice)
	err = request(ctx, c.api, req, &result)
	return
}

// SingleShortTakeProfitMarket 单向持仓 看空仓位 市价止盈
func (c custom) SingleShortTakeProfitMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = new(trade.PlaceData)
	req.Params.SingleShortTakeProfitMarket(symbol, quantity, stopPrice)
	err = request(ctx, c.api, req, &result)
	return
}

// SingleLongStopLossMarket 单向持仓 看多仓位 市价止损
func (c custom) SingleLongStopLossMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = new(trade.PlaceData)
	req.Params.SingleLongStopLossMarket(symbol, quantity, stopPrice)
	err = request(ctx, c.api, req, &result)
	return
}

// SingleShortStopLossMarket 单向持仓 看空仓位 市价止损
func (c custom) SingleShortStopLossMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	req.Params = new(trade.PlaceData)
	req.Params.SingleShortStopLossMarket(symbol, quantity, stopPrice)
	err = request(ctx, c.api, req, &result)
	return
}
