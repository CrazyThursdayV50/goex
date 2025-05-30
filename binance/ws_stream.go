package binance

import (
	"github.com/CrazyThursdayV50/goex/binance/websocket-streams/stream"
)

type WebSocketStreams = stream.WebSocketStreams

func NewWebSocketStreams() *WebSocketStreams {
	return &WebSocketStreams{}
}
