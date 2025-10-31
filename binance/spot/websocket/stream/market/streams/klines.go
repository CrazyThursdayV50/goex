package streams

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models/kline"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/spot"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
)

func (s *Stream) KlinesStream(streamName string, handler WsHandlerKlines) *Stream {
	stream := s.Clone()
	wsclient := client.NewClient(
		stream.logger,
		fmt.Sprintf(spot.WsStream().Endpoint(), streamName),
		variables.GetProxy(),
		stream.handler,
		nil,
	)
	stream.wsclient = wsclient

	stream.HandleKlineEvent(handler)
	return stream
}

// for klines stream
func (s *Stream) HandleKlineEvent(fn func(*kline.Data, error)) {
	s.RegisterEventHandler(
		kline.Event,
		CreateBytesHandler(fn),
	)
}

// for klines combined
func (s *Stream) HandleKlineStream(symbol, interval, timezone string, fn func(*kline.Data, error)) {
	s.RegisterEventHandler(
		kline.StreamName(symbol, interval, timezone),
		CreateBytesHandler(fn),
	)
}
