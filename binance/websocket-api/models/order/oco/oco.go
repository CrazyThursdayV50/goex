package oco

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models"
	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models/order"
)

type ParamsData struct {
	Symbol            string          `json:"symbol"`
	ListClientOrderID *string         `json:"listClientOrderId,omitempty"`
	Side              order.OrderSide `json:"side"`
	Quantity          string          `json:"quantity"`

	AboveType           order.OrderType      `json:"aboveType"`
	AboveClientOrderID  *string              `json:"aboveClientOrderId,omitempty"`
	AboveIcebergQty     *string              `json:"aboveIcebergQty,omitempty"`
	AbovePrice          *string              `json:"abovePrice,omitempty"`
	AboveStopPrice      *string              `json:"aboveStopPrice,omitempty"`
	AboveTrailingDelta  *string              `json:"aboveTrailingDelta,omitempty"`
	AboveTimeInForce    *order.TimeInForce   `json:"aboveTimeInForce,omitempty"`
	AboveStrategyID     *string              `json:"aboveStrategyID,omitempty"`
	AboveStrategyType   *int64               `json:"aboveStrategyType,omitempty"`
	AbovePegPriceType   *order.PegPriceType  `json:"abovePegPriceType,omitempty"`
	AbovePegOffsetType  *order.PegOffsetType `json:"abovePegOffsetType,omitempty"`
	AbovePegOffsetValue *int64               `json:"abovePegOffsetValue,omitempty"`

	BelowType           order.OrderType      `json:"belowType"`
	BelowClientOrderID  *string              `json:"belowClientOrderId,omitempty"`
	BelowIcebergQty     *string              `json:"belowIcebergQty,omitempty"`
	BelowPrice          *string              `json:"belowPrice,omitempty"`
	BelowStopPrice      *string              `json:"belowStopPrice,omitempty"`
	BelowTrailingDelta  *string              `json:"belowTrailingDelta,omitempty"`
	BelowTimeInForce    *order.TimeInForce   `json:"belowTimeInForce,omitempty"`
	BelowStrategyID     *string              `json:"belowStrategyID,omitempty"`
	BelowStrategyType   *int64               `json:"belowStrategyType,omitempty"`
	BelowPegPriceType   *order.PegPriceType  `json:"belowPegPriceType,omitempty"`
	BelowPegOffsetType  *order.PegOffsetType `json:"belowPegOffsetType,omitempty"`
	BelowPegOffsetValue *int64               `json:"belowPegOffsetValue,omitempty"`

	NewOrderRespType        *order.OrderRespType           `json:"newOrderRespType,omitempty"`
	SelfTradePreventionMode *order.SelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`

	models.Sign
}

func (p *ParamsData) Map() map[string]string {
	params := p.Sign.Map()
	params["symbol"] = p.Symbol
	params["side"] = string(p.Side)
	params["quantity"] = p.Quantity
	params["aboveType"] = string(p.AboveType)
	params["belowType"] = string(p.BelowType)

	if p.ListClientOrderID != nil {
		params["listClientOrderId"] = *p.ListClientOrderID
	}

	if p.AboveClientOrderID != nil {
		params["aboveClientOrderId"] = *p.AboveClientOrderID
	}

	if p.BelowClientOrderID != nil {
		params["belowClientOrderId"] = *p.BelowClientOrderID
	}

	if p.AboveIcebergQty != nil {
		params["aboveIcebergQty"] = *p.AboveIcebergQty
	}

	if p.BelowIcebergQty != nil {
		params["belowIcebergQty"] = *p.BelowIcebergQty
	}

	if p.AbovePrice != nil {
		params["abovePrice"] = *p.AbovePrice
	}

	if p.BelowPrice != nil {
		params["belowPrice"] = *p.BelowPrice
	}

	if p.AboveStopPrice != nil {
		params["aboveStopPrice"] = *p.AboveStopPrice
	}

	if p.BelowStopPrice != nil {
		params["belowStopPrice"] = *p.BelowStopPrice
	}

	if p.AboveTrailingDelta != nil {
		params["aboveTrailingDelta"] = *p.AboveTrailingDelta
	}

	if p.BelowTrailingDelta != nil {
		params["belowTrailingDelta"] = *p.BelowTrailingDelta
	}

	if p.AboveTimeInForce != nil {
		params["aboveTimeInForce"] = string(*p.AboveTimeInForce)
	}

	if p.BelowTimeInForce != nil {
		params["belowTimeInForce"] = string(*p.BelowTimeInForce)
	}

	if p.AboveStrategyID != nil {
		params["aboveStrategyID"] = *p.AboveStrategyID
	}

	if p.BelowStrategyID != nil {
		params["belowStrategyID"] = *p.BelowStrategyID
	}

	if p.AboveStrategyType != nil {
		params["aboveStrategyType"] = string(*p.AboveStrategyType)
	}

	if p.BelowStrategyType != nil {
		params["belowStrategyType"] = string(*p.BelowStrategyType)
	}

	if p.AbovePegPriceType != nil {
		params["abovePegPriceType"] = string(*p.AbovePegPriceType)
	}

	if p.BelowPegPriceType != nil {
		params["belowPegPriceType"] = string(*p.BelowPegPriceType)
	}

	if p.AbovePegOffsetType != nil {
		params["abovePegOffsetType"] = string(*p.AbovePegOffsetType)
	}

	if p.BelowPegOffsetType != nil {
		params["belowPegOffsetType"] = string(*p.BelowPegOffsetType)
	}

	if p.AbovePegOffsetValue != nil {
		params["abovePegOffsetValue"] = fmt.Sprintf("%d", *p.AbovePegOffsetValue)
	}

	if p.BelowPegOffsetValue != nil {
		params["belowPegOffsetValue"] = fmt.Sprintf("%d", *p.BelowPegOffsetValue)
	}

	if p.NewOrderRespType != nil {
		params["newOrderRespType"] = string(*p.NewOrderRespType)
	}

	if p.SelfTradePreventionMode != nil {
		params["selfTradePreventionMode"] = string(*p.SelfTradePreventionMode)
	}

	return params
}

// Params WebSocket API 下OCO单请求
type Params = models.WsAPIParams[*ParamsData]

func NewParams() *Params {
	return &Params{
		Method: "orderList.place.oco",
	}
}

type ResultData struct {
	OrderListID       int64                 `json:"orderListId"`
	ContingencyType   string                `json:"contingencyType"`
	ListStatusType    order.ListStatusType  `json:"listStatusType"`
	ListOrderStatus   order.ListOrderStatus `json:"listOrderStatus"`
	ListClientOrderID string                `json:"listClientOrderId"`
	TransactionTime   int64                 `json:"transactionTime"`
	Symbol            string                `json:"symbol"`
	OrderReports      []*order.ResultData   `json:"orderReports"`
	Orders            []*order.ResultData   `json:"orders"`
}
