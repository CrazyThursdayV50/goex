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
	req.Params, err = t.api.Sign(data)
	if err != nil {
		return
	}

	err = request(ctx, t.api.api, req, &result)
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
	var data trade.PlaceData
	data.SingleOpenLongMarket(symbol, quantity)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// OpenShortMarketSingle 单向持仓市价开空
func (c custom) SingleOpenShortMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.SingleOpenShortMarket(symbol, quantity)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// ReduceLongMarketSingle 单向持仓 看多仓位 市价减仓
func (c custom) SingleReduceLongMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.SingleReduceLongMarket(symbol, quantity)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}

	err = request(ctx, c.api.api, req, &result)
	return
}

// ReduceShortMarketSingle 单向持仓 看空仓位 市价减仓
func (c custom) SingleReduceShortMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.SingleReduceShortMarket(symbol, quantity)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}

	err = request(ctx, c.api.api, req, &result)
	return
}

// SingleLongTakeProfitMarket 单向持仓 看多仓位 市价止盈
func (c custom) SingleLongTakeProfitMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.SingleLongTakeProfitMarket(symbol, quantity, stopPrice)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// SingleShortTakeProfitMarket 单向持仓 看空仓位 市价止盈
func (c custom) SingleShortTakeProfitMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.SingleShortTakeProfitMarket(symbol, quantity, stopPrice)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// SingleLongStopLossMarket 单向持仓 看多仓位 市价止损
func (c custom) SingleLongStopLossMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.SingleLongStopLossMarket(symbol, quantity, stopPrice)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// SingleShortStopLossMarket 单向持仓 看空仓位 市价止损
func (c custom) SingleShortStopLossMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.SingleShortStopLossMarket(symbol, quantity, stopPrice)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}
