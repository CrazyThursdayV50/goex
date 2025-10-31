package stream

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/streams"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type Stream struct{}

func New() Stream {
	return Stream{}
}

func (Stream) Market(logger log.Logger) *streams.Stream {
	return streams.New(logger)
}
