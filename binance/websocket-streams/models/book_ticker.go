package models

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// IndividualSymbolBookTickerEvent 个股最佳买卖价事件
/*
	{
	  "u":400900217,     // order book updateId
	  "s":"BNBUSDT",     // symbol
	  "b":"25.35190000", // best bid price
	  "B":"31.21000000", // best bid qty
	  "a":"25.36520000", // best ask price
	  "A":"40.66000000"  // best ask qty
	}
*/
type IndividualSymbolBookTickerEvent = CombinedEvent[*IndividualSymbolBookTicker]

// IndividualSymbolBookTicker 个股最佳买卖价数据
type IndividualSymbolBookTicker struct {
	UpdateId    int64           `json:"u"`
	Symbol      string          `json:"s"`
	BidPrice    decimal.Decimal `json:"b"`
	BidQuantity decimal.Decimal `json:"B"`
	AskPrice    decimal.Decimal `json:"a"`
	AskQuantity decimal.Decimal `json:"A"`
}

func (e *IndividualSymbolBookTicker) String() string {
	if e == nil {
		return "nil"
	}

	return fmt.Sprintf("%s - [%d]ask: %s@%s, bid: %s@%s", e.Symbol, e.UpdateId, e.AskQuantity, e.AskPrice, e.BidQuantity, e.BidPrice)
} 