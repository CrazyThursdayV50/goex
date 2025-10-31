package binance

import (
	"github.com/CrazyThursdayV50/goex/binance/spot"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type SpotEntry struct{}

func (SpotEntry) WebSocketStream() *spot.WebSocketStreams {
	return spot.NewWebSocketStreams()
}

func (SpotEntry) NewWebSocketAPI(logger log.Logger, apiKey, secretKey string) *spot.WebsocketAPI {
	return spot.NewWebSocketAPI(logger, apiKey, secretKey)
}

func Spot() SpotEntry { return SpotEntry{} }
