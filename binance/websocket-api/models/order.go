package models

import (
	"fmt"
)

// OrderSide 订单方向
type OrderSide string

const (
	BUY  OrderSide = "BUY"
	SELL OrderSide = "SELL"
)

// OrderType 订单类型
type OrderType string

const (
	LIMIT             OrderType = "LIMIT"
	MARKET            OrderType = "MARKET"
	STOP_LOSS         OrderType = "STOP_LOSS"
	STOP_LOSS_LIMIT   OrderType = "STOP_LOSS_LIMIT"
	TAKE_PROFIT       OrderType = "TAKE_PROFIT"
	TAKE_PROFIT_LIMIT OrderType = "TAKE_PROFIT_LIMIT"
	LIMIT_MAKER       OrderType = "LIMIT_MAKER"
)

// TimeInForce 订单有效期
type TimeInForce string

const (
	GTC TimeInForce = "GTC" // Good Till Cancel
	IOC TimeInForce = "IOC" // Immediate or Cancel
	FOK TimeInForce = "FOK" // Fill or Kill
)

// WsOrderParamsData WebSocket API 下单请求参数
type WsOrderParamsData struct {
	Symbol                  string       `json:"symbol"`                            // 交易对
	Side                    OrderSide    `json:"side"`                              // 订单方向
	Type                    OrderType    `json:"type"`                              // 订单类型
	TimeInForce             *TimeInForce `json:"timeInForce,omitempty"`             // 订单有效期
	Price                   *string      `json:"price,omitempty"`                   // 价格
	Quantity                *string      `json:"quantity,omitempty"`                // 数量
	QuoteOrderQty           *string      `json:"quoteOrderQty,omitempty"`           // 报价数量
	NewClientOrderId        *string      `json:"newClientOrderId,omitempty"`        // 客户端订单ID
	StopPrice               *string      `json:"stopPrice,omitempty"`               // 止损价
	IcebergQty              *string      `json:"icebergQty,omitempty"`              // 冰山数量
	NewOrderRespType        *string      `json:"newOrderRespType,omitempty"`        // 响应类型
	SelfTradePreventionMode *string      `json:"selfTradePreventionMode,omitempty"` // 自成交预防模式
	StrategyId              *int64       `json:"strategyId,omitempty"`              // 策略ID
	StrategyType            *int64       `json:"strategyType,omitempty"`            // 策略类型
	TrailingDelta           *int64       `json:"trailingDelta,omitempty"`           // 追踪止损点数
	Sign
}

// Map 返回用于签名的参数map
func (p *WsOrderParamsData) Map() map[string]string {
	params := p.Sign.Map()
	params["symbol"] = p.Symbol
	params["side"] = string(p.Side)
	params["type"] = string(p.Type)

	if p.TimeInForce != nil {
		params["timeInForce"] = string(*p.TimeInForce)
	}

	if p.Price != nil {
		params["price"] = *p.Price
	}
	if p.Quantity != nil {
		params["quantity"] = *p.Quantity
	}
	if p.QuoteOrderQty != nil {
		params["quoteOrderQty"] = *p.QuoteOrderQty
	}
	if p.NewClientOrderId != nil {
		params["newClientOrderId"] = *p.NewClientOrderId
	}
	if p.StopPrice != nil {
		params["stopPrice"] = *p.StopPrice
	}
	if p.IcebergQty != nil {
		params["icebergQty"] = *p.IcebergQty
	}
	if p.NewOrderRespType != nil {
		params["newOrderRespType"] = *p.NewOrderRespType
	}
	if p.SelfTradePreventionMode != nil {
		params["selfTradePreventionMode"] = *p.SelfTradePreventionMode
	}
	if p.StrategyId != nil {
		params["strategyId"] = fmt.Sprintf("%d", *p.StrategyId)
	}
	if p.StrategyType != nil {
		params["strategyType"] = fmt.Sprintf("%d", *p.StrategyType)
	}
	if p.TrailingDelta != nil {
		params["trailingDelta"] = fmt.Sprintf("%d", *p.TrailingDelta)
	}

	return params
}

// WsOrderParams WebSocket API 下单请求
type WsOrderParams = WsAPIParams[*WsOrderParamsData]

func NewWsOrderParams() *WsOrderParams {
	return &WsOrderParams{
		Method: "order.place",
	}
}

// WsOrderResultData WebSocket API 下单响应数据
type WsOrderResultData struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactTime            int64  `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             int64  `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	StopPrice               string `json:"stopPrice,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           int64  `json:"trailingDelta,omitempty"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
}

// WsTestOrderParams WebSocket API 测试下单请求
type WsTestOrderParams = WsAPIParams[*WsOrderParamsData]

func NewWsTestOrderParams() *WsTestOrderParams {
	return &WsTestOrderParams{
		Method: "order.test",
	}
} 