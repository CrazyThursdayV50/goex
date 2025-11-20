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
type single struct {
	api *API
}

func (t *Trade) Single() single {
	return single{api: t.api}
}

// OpenLongMarketSingle 市价开多
func (c single) OpenLongMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
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

// OpenShortMarketSingle 市价开空
func (c single) OpenShortMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
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

// ReduceLongMarketSingle 看多仓位 市价减仓
func (c single) ReduceLongMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
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

// ReduceShortMarketSingle 看空仓位 市价减仓
func (c single) ReduceShortMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
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

// LongTakeProfitMarket 看多仓位 市价止盈
func (c single) LongTakeProfitMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
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

// ShortTakeProfitMarket 看空仓位 市价止盈
func (c single) ShortTakeProfitMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
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

// LongStopLossMarket 看多仓位 市价止损
func (c single) LongStopLossMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
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

// ShortStopLossMarket 看空仓位 市价止损
func (c single) ShortStopLossMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
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

// 双向持仓
type dual struct {
	api *API
}

func (t *Trade) Dual() dual { return dual{api: t.api} }

// OpenLongMarketSingle 市价开多
func (c dual) OpenLongMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.DualOpenLongMarket(symbol, quantity)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// OpenShortMarketSingle 市价开空
func (c dual) OpenShortMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.DualOpenShortMarket(symbol, quantity)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// ReduceLongMarketSingle 看多仓位 市价减仓
func (c dual) ReduceLongMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.DualReduceLongMarket(symbol, quantity)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}

	err = request(ctx, c.api.api, req, &result)
	return
}

// ReduceShortMarketSingle 看空仓位 市价减仓
func (c dual) ReduceShortMarket(ctx context.Context, symbol, quantity string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.DualReduceShortMarket(symbol, quantity)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}

	err = request(ctx, c.api.api, req, &result)
	return
}

// LongTakeProfitMarket 看多仓位 市价止盈
func (c dual) LongTakeProfitMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.DualLongTakeProfitMarket(symbol, quantity, stopPrice)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// ShortTakeProfitMarket 看空仓位 市价止盈
func (c dual) ShortTakeProfitMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.DualShortTakeProfitMarket(symbol, quantity, stopPrice)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// LongStopLossMarket 看多仓位 市价止损
func (c dual) LongStopLossMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.DualLongStopLossMarket(symbol, quantity, stopPrice)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}

// ShortStopLossMarket 看空仓位 市价止损
func (c dual) ShortStopLossMarket(ctx context.Context, symbol, quantity, stopPrice string) (result *trade.PlaceResultData, err error) {
	req := trade.Place()
	var data trade.PlaceData
	data.DualShortStopLossMarket(symbol, quantity, stopPrice)
	req.Params, err = c.api.Sign(&data)
	if err != nil {
		return
	}
	err = request(ctx, c.api.api, req, &result)
	return
}
