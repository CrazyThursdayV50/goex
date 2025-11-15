// 标记价格
package streams

import (
	"fmt"

	markprice "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/markPrice"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
)

func (s *Stream) MarkPriceStream(streamName string, handler func(*markprice.Result, error)) *Stream {
	stream := s.Clone()
	stream.wsclient = client.NewClient(
		stream.logger,
		fmt.Sprintf("%s/%s", derivatives.WsStream().Endpoint(), streamName),
		variables.GetProxy(),
		stream.handler,
		nil,
	)

	stream.HandleMarkPriceStreamEvent(handler)
	return stream
}

func (s *Stream) HandleMarkPriceStreamEvent(f func(*markprice.Result, error)) *Stream {
	s.RegisterEventHandler(
		markprice.Event,
		CreateBytesHandler(f),
	)
	return s
}

func (s *Stream) HandleMarkPriceCombinedData(streamName string, f func(*markprice.Result, error)) *Stream {
	s.RegisterEventHandler(
		streamName,
		CreateBytesHandler(f),
	)
	return s
}
