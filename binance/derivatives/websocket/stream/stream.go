package stream

import (
	market "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/streams"
	user "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/streams"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type Stream struct{}

func New() Stream {
	return Stream{}
}

func (Stream) Market(logger log.Logger) *market.Stream {
	return market.New(logger)
}

func (Stream) User(logger log.Logger, listenKey string) *user.Stream {
	return user.New(logger, listenKey)
}
