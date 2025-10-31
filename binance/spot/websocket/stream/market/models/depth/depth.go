package depth

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/pkgo/builtin/collector"
	"github.com/shopspring/decimal"
)

func StreamName(symbol string, level int) string {
	return fmt.Sprintf("%s@depth%d", strings.ToLower(symbol), level)
}

// PartialDepthEvent 深度数据事件
//
//	{
//	  "lastUpdateId": 160,  // Last update ID
//	  "bids": [             // Bids to be updated
//	    [
//	      "0.0024",         // Price level to be updated
//	      "10"              // Quantity
//	    ]
//	  ],
//	  "asks": [             // Asks to be updated
//	    [
//	      "0.0026",         // Price level to be updated
//	      "100"             // Quantity
//	    ]
//	  ]
//	}
type PartialDepthEvent struct {
	LastupdateID int64               `json:"lastUpdateId"`
	Bids         [][]decimal.Decimal `json:"bids"`
	Asks         [][]decimal.Decimal `json:"asks"`
}

func (m *PartialDepthEvent) PartialDepthData() *PartialDepthData {
	var data PartialDepthData
	data.LastupdateId = m.LastupdateID
	data.Bids = collector.Slice(m.Bids, func(i int, e []decimal.Decimal) (bool, *orderbook) {
		orderbook := newOrderBook(e)
		return orderbook != nil, orderbook
	})

	data.Asks = collector.Slice(m.Asks, func(i int, e []decimal.Decimal) (bool, *orderbook) {
		orderbook := newOrderBook(e)
		return orderbook != nil, orderbook
	})

	return &data
}

// orderbook 订单簿条目
type orderbook struct {
	Price    decimal.Decimal
	Quantity decimal.Decimal
}

func (o *orderbook) String() string {
	if o == nil {
		return "nil"
	}

	return fmt.Sprintf("%s@%s", o.Quantity.String(), o.Price.String())
}

func newOrderBook(sli []decimal.Decimal) *orderbook {
	if len(sli) != 2 {
		return nil
	}

	var b orderbook
	b.Price = sli[0]
	b.Quantity = sli[1]
	return &b
}

// PartialDepthData 深度数据
type PartialDepthData struct {
	LastupdateId int64
	Bids         []*orderbook
	Asks         []*orderbook
}

func (d *PartialDepthData) String() string {
	if d == nil {
		return "nil"
	}

	return fmt.Sprintf("[%d]asks: %s, bids: %s", d.LastupdateId, d.Asks, d.Bids)
}

// CombinedEvent 组合流事件通用结构
// {"stream":"bnbusdt@depth5@100ms","data":{"lastUpdateId":13919791373,"bids":[["656.72000000","29.05200000"],["656.71000000","11.68500000"],["656.70000000","1.01600000"],["656.69000000","0.00800000"],["656.68000000","0.02500000"]],"asks":[["656.73000000","21.01200000"],["656.74000000","0.02400000"],["656.75000000","0.02400000"],["656.76000000","0.17100000"],["656.77000000","0.62500000"]]}}
type CombinedEvent[T any] struct {
	Stream string `json:"stream"`
	Data   T      `json:"data"`
}

// PartialDepthCombinedEvent 组合深度数据事件
type PartialDepthCombinedEvent CombinedEvent[*PartialDepthEvent]

// PartialDepthCombinedData 组合深度数据
type PartialDepthCombinedData struct {
	Symbol string
	PartialDepthData
}

func (d *PartialDepthCombinedData) String() string {
	if d == nil {
		return "nil"
	}

	return fmt.Sprintf("%s - %s", d.Symbol, d.PartialDepthData.String())
}

func parsePartialDepth(stream string) string {
	sli := strings.Split(stream, "@")
	switch len(sli) {
	case 2, 3:
		return sli[0]
	default:
		return ""
	}
}

func (e *PartialDepthCombinedEvent) PartialDepthCombinedData() *PartialDepthCombinedData {
	var data PartialDepthCombinedData
	data.Symbol = strings.ToUpper(parsePartialDepth(e.Stream))
	if data.Symbol == "" {
		return nil
	}

	data.PartialDepthData = *e.Data.PartialDepthData()
	return &data
}
