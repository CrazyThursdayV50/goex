package stream

import (
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/streams"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type Stream struct {
}

func (Stream) Market(logger log.Logger) *streams.Stream {
	return streams.New(logger)
}

func New() Stream {
	return Stream{}
}
