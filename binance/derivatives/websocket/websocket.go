package websocket

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream"
)

type Websocket struct{}

func New() Websocket {
	return Websocket{}
}

func (Websocket) API() api.API {
	return api.New()
}

func (Websocket) Stream() stream.Stream {
	return stream.New()
}
