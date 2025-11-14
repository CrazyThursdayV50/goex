package continuous

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/klines"
)

const (
	Event = "continuous_kline"
	// <pair>_<contractType>@continuousKline_<interval>
	streamNameFormat = "%s_%s@continuousKline_%s"
)

func StreamName(symbol, contractType, interval string) string {
	return fmt.Sprintf(streamNameFormat, strings.ToLower(symbol), contractType, interval)
}

type Result struct {
	models.BaseResult
	Symbol   string           `json:"ps"`
	Contract string           `json:"ct"`
	Data     klines.KlineData `json:"k"`
}
