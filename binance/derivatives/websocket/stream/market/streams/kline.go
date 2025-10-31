package streams

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/kline"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/kline/continuous"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
)

func (s *Stream) KlinesStream(streamName string, handler func(*kline.KlineData, error)) *Stream {
	stream := s.Clone()
	stream.wsclient = client.NewClient(
		stream.logger,
		fmt.Sprintf("%s/%s", derivatives.WsStream().Endpoint(), streamName),
		variables.GetProxy(),
		stream.handler,
		nil,
	)

	stream.HandleKlineStreamEvent(handler)
	return stream
}

func (s *Stream) ContinuousKlineStream(streamName string, handler func(*kline.KlineData, error)) *Stream {
	return s.KlinesStream(streamName, handler)
}

func (s *Stream) HandleKlineStreamEvent(f func(*kline.KlineData, error)) *Stream {
	s.RegisterEventHandler(
		kline.Event,
		CreateBytesHandler(f),
	)
	return s
}

func (s *Stream) HandleKlineCombinedData(streamName string, f func(*kline.KlineData, error)) *Stream {
	s.RegisterEventHandler(
		streamName,
		CreateBytesHandler(f),
	)
	return s
}

func (s *Stream) HandleContinuousKlineStreamEvent(f func(*kline.KlineData, error)) *Stream {
	s.RegisterEventHandler(
		continuous.Event,
		CreateBytesHandler(f),
	)
	return s
}

func (s *Stream) HandleContinuousKlineCombinedData(streamName string, f func(*kline.KlineData, error)) *Stream {
	s.RegisterEventHandler(
		streamName,
		CreateBytesHandler(f),
	)
	return s
}
