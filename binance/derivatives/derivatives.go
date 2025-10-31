package derivatives

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream"
)

type Derivatives struct {
}

func (Derivatives) WebSocketAPI() api.API {
	return api.New()
}

func (Derivatives) WebSocketStream() stream.Stream {
	return stream.New()
}
