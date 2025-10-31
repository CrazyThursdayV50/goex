package spot

import (
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-streams/stream"
)

type WebSocketStreams = stream.WebSocketStreams

func NewWebSocketStreams() *WebSocketStreams {
	return &WebSocketStreams{}
}
