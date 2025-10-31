package trade

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"
	"github.com/CrazyThursdayV50/goex/infra/utils"
)

// 参数理解：
// 双向持仓（估计下面说的“双开”也是这个意思）时，有四种操作模式：开多、开空、平多、平空。
// 需要用 OrderSide(side) + PositionSide 两个参数共同来指定
// 单向持仓时，OrderSide 的买入和卖出分别对应开多和开空，当 ReduceOnly 为 true 时，OrderSide 的买入和卖出分别对应平空和平多
// closePosition 指当前持仓全部平仓。
type PlaceData struct {
	Symbol    string    `json:"symbol"`
	OrderSide OrderSide `json:"side"`

	// 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	PositionSide *PositionSide `json:"positionSide,omitempty"`

	OrderType OrderType `json:"type,omitempty"`
	// true, false; 非双开模式下默认false；双开模式下不接受此参数； 使用closePosition不支持此参数。
	ReduceOnly *ReduceOnly `json:"reduceOnly,omitempty"`
	// 下单数量,使用closePosition不支持此参数。
	Quantity         *string `json:"quantity,omitempty"`
	Price            *string `json:"price,omitempty"`
	NewClientOrderID *string `json:"newClientOrderId,omitempty"`

	// 止盈止损触发价格
	StopPrice *string `json:"stopPrice,omitempty"`
	// true, false；触发后全部平仓，仅支持STOP_MARKET和TAKE_PROFIT_MARKET；不与quantity合用；自带只平仓效果，不与reduceOnly 合用
	ClosePosition *ClosePosition `json:"closePosition,omitempty"`

	// 追踪止损激活价格，仅 TRAILING_STOP_MARKET 需要此参数, 默认为下单当前市场价格(支持不同workingType)
	ActivationPrice *string `json:"activationPrice,omitempty"`
	// 追踪止损回调比例，可取值范围[0.1, 10],其中 1代表1% ,仅TRAILING_STOP_MARKET 需要此参数
	CallbackRate *string      `json:"callbackRate,omitempty"`
	TimeInForce  *TimeInForce `json:"timeInForce,omitempty"`
	// stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	WorkingType *WorkingType `json:"workingType,omitempty"`
	// 条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	PriceProtect     *PriceProtect     `json:"priceProtect,omitempty"`
	NewOrderRespType *NewOrderRespType `json:"newOrderRespType,omitempty"`

	// OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	PriceMatch              *PriceMatch              `json:"priceMatch,omitempty"`
	SelfTradePreventionMode *SelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`
	GoodTillDate            *int64                   `json:"goodTillDate,omitempty"`
}

func (d *PlaceData) Base(symbol string) {
	d.Symbol = symbol
	d.NewOrderRespType = utils.Ptr(NEW_ORDER_RESP_TYPE_RESULT)
}

func (d *PlaceData) SingleOpenLongMarket(symbol, quantity string) {
	d.Base(symbol)
	d.PositionSide = nil
	d.OrderSide = SIDE_BUY
	d.OrderType = TYPE_MARKET
	d.Quantity = utils.Ptr(quantity)
	d.ReduceOnly = nil
}

func (d *PlaceData) SingleOpenShortMarket(symbol, quantity string) {
	d.Base(symbol)
	d.PositionSide = nil
	d.OrderSide = SIDE_SELL
	d.OrderType = TYPE_MARKET
	d.Quantity = utils.Ptr(quantity)
	d.ReduceOnly = nil
}

// 单向持仓 多仓 减仓
func (d *PlaceData) SingleReduceLongMarket(symbol, quantity string) {
	d.SingleOpenShortMarket(symbol, quantity)
	d.ReduceOnly = utils.Ptr(REDUCE_ONLY_TRUE)
}

// 单向持仓 空仓 减仓
func (d *PlaceData) SingleReduceShortMarket(symbol, quantity string) {
	d.SingleOpenLongMarket(symbol, quantity)
	d.ReduceOnly = utils.Ptr(REDUCE_ONLY_TRUE)
}

// 单向持仓 多仓 止盈
func (d *PlaceData) SingleLongTakeProfitMarket(symbol, quantity, stopPrice string) *PlaceData {
	d.SingleReduceLongMarket(symbol, quantity)
	d.OrderType = TYPE_TAKE_PROFIT_MARKET
	d.StopPrice = utils.Ptr(stopPrice)
	return d
}

// 单向持仓 多仓 止损
func (d *PlaceData) SingleLongStopLossMarket(symbol, quantity, stopPrice string) *PlaceData {
	d.SingleReduceLongMarket(symbol, quantity)
	d.OrderType = TYPE_STOP_MARKET
	d.StopPrice = utils.Ptr(stopPrice)
	return d
}

// 单向持仓 空仓 止盈
func (d *PlaceData) SingleShortTakeProfitMarket(symbol, quantity, stopPrice string) *PlaceData {
	d.SingleReduceShortMarket(symbol, quantity)
	d.OrderType = TYPE_TAKE_PROFIT_MARKET
	d.StopPrice = utils.Ptr(stopPrice)
	return d
}

// 单向持仓 空仓 止损
func (d *PlaceData) SingleShortStopLossMarket(symbol, quantity, stopPrice string) *PlaceData {
	d.SingleReduceShortMarket(symbol, quantity)
	d.OrderType = TYPE_STOP_MARKET
	d.StopPrice = utils.Ptr(stopPrice)
	return d
}

// TODO: 双向持仓
// var _ iface.SignerData = (*PlaceData)(nil)

func Place() *models.Request {
	return &models.Request{
		Method: "order.place",
	}
}

// func PlaceTest() *PlaceRequest {
// 	return &PlaceRequest{
// 		Method: "order.test",
// 	}
// }

/*
	{
	  "orderId": 325078477,
	  "symbol": "BTCUSDT",
	  "status": "NEW",
	  "clientOrderId": "iCXL1BywlBaf2sesNUrVl3",
	  "price": "43187.00",
	  "avgPrice": "0.00",
	  "origQty": "0.100",
	  "executedQty": "0.000",
	  "cumQty": "0.000",
	  "cumQuote": "0.00000",
	  "timeInForce": "GTC",
	  "type": "LIMIT",
	  "reduceOnly": false,
	  "closePosition": false,
	  "side": "BUY",
	  "positionSide": "BOTH",
	  "stopPrice": "0.00",
	  "workingType": "CONTRACT_PRICE",
	  "priceProtect": false,
	  "origType": "LIMIT",
	  "priceMatch": "NONE",
	  "selfTradePreventionMode": "NONE",
	  "goodTillDate": 0,
	  "updateTime": 1702555534435
	}
*/
type PlaceResultData struct {
	OrderID                 int64                   `json:"orderId"`
	Symbol                  string                  `json:"symbol"`
	Status                  OrderStatus             `json:"status"`
	ClientOrderID           string                  `json:"clientOrderId"`
	Price                   string                  `json:"price"`
	AvgPrice                string                  `json:"avgPrice"`
	OrigQty                 string                  `json:"origQty"`
	ExecutedQty             string                  `json:"executedQty"`
	CumQty                  string                  `json:"cumQty"`
	CumQuote                string                  `json:"cumQuote"`
	TimeInForce             TimeInForce             `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	ClosePosition           bool                    `json:"closePosition"`
	Side                    OrderSide               `json:"side"`
	PositionSide            PositionSide            `json:"positionSide"`
	StopPrice               string                  `json:"stopPrice"`
	WorkingType             WorkingType             `json:"workingType"`
	PriceProtect            bool                    `json:"priceProtect"`
	OrigType                OrderType               `json:"origType"`
	PriceMatch              PriceMatch              `json:"priceMatch"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	GoodTillDate            int64                   `json:"goodTillDate"`
	UpdateTime              int64                   `json:"updateTime"`
}
