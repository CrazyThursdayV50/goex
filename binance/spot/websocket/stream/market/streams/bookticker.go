// 最佳买卖数据
package streams

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models/bookticker"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/spot"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
)

func (s *Stream) IndividualBookTicker(streamName string, handler WsIndividualSymbolBookTickerHandler) *Stream {
	stream := s.Clone()
	wsclient := client.NewClient(
		stream.logger,
		fmt.Sprintf(spot.WsStream().Endpoint(), streamName),
		variables.GetProxy(),
		stream.handler,
		nil,
	)

	stream.wsclient = wsclient
	stream.HandleUnexpected(CreateBytesHandler(handler))
	return stream
}

// for combined stream
func (s *Stream) HandleBookTicker(
	streamName string,
	fn func(*bookticker.IndividualSymbolBookTicker, error),
) {
	s.RegisterEventHandler(
		streamName,
		CreateBytesHandler(fn),
	)
}
