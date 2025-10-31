package spot

import (
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream"
)

type SpotEntry struct {
}

func New() SpotEntry {
	return SpotEntry{}
}

func (SpotEntry) WebsocketStream() stream.Stream {
	return stream.New()
}

func (SpotEntry) WebSocketAPI() api.API {
	return api.New()
}
