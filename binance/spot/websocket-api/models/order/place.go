// Package order
package order

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models"
)

// ParamsData WebSocket API 下单请求参数
type ParamsData struct {
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
	models.Sign
}

// Map 返回用于签名的参数map
func (p *ParamsData) Map() map[string]string {
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

// Params WebSocket API 下单请求
type Params = models.WsAPIParams[*ParamsData]

func NewParams() *Params {
	return &Params{
		Method: "order.place",
	}
}

func NewParamsTest() *Params {
	return &Params{
		Method: "order.test",
	}
}

// ResultData WebSocket API 下单响应数据
type ResultData struct {
	Symbol                  string                  `json:"symbol"`
	OrderID                 int64                   `json:"orderId"`
	OrderListID             int64                   `json:"orderListId"`
	ClientOrderID           string                  `json:"clientOrderId"`
	TransactTime            int64                   `json:"transactTime"`
	Price                   string                  `json:"price"`
	OrigQty                 string                  `json:"origQty"`
	ExecutedQty             string                  `json:"executedQty"`
	OrigQuoteOrderQty       string                  `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     string                  `json:"cummulativeQuoteQty"`
	Status                  OrderStatus             `json:"status"`
	TimeInForce             TimeInForce             `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	Side                    OrderSide               `json:"side"`
	WorkingTime             int64                   `json:"workingTime"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	Fills                   []*Trade                `json:"fills,omitempty"`

	// 特殊情况下才有的字段
	IcebergQty        string        `json:"icebergQty,omitempty"`
	PreventedMatchID  int64         `json:"preventedMatchId,omitempty"`
	PreventedQuantity string        `json:"preventedQuantity,omitempty"`
	StopPrice         string        `json:"stopPrice,omitempty"`
	StrategyID        int64         `json:"strategyId,omitempty"`
	StrategyType      int64         `json:"strategyType,omitempty"`
	TrailingDelta     int64         `json:"trailingDelta,omitempty"`
	TrailingTime      int64         `json:"trailingTime,omitempty"`
	UsedSor           bool          `json:"usedSor,omitempty"`
	WorkingFloor      string        `json:"workingFloor,omitempty"`
	PegPriceType      PegPriceType  `json:"pegPriceType,omitempty"`
	PegOffsetType     PegOffsetType `json:"pegOffsetType,omitempty"`
	PegOffsetValue    int64         `json:"pegOffsetValue,omitempty"`
	PeggedPrice       string        `json:"peggedPrice,omitempty"`
}
