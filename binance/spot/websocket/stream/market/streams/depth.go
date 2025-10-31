// 有限档深度信息
package streams

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models/depth"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/spot"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
)

func (s *Stream) PartialBookDepth5Stream(streamName string, handler WsPartialDepthHandler) *Stream {
	stream := s.Clone()
	wsclient := client.NewClient(
		stream.logger,
		fmt.Sprintf(spot.WsStream().Endpoint(), streamName),
		variables.GetProxy(),
		stream.handler,
		nil,
	)
	stream.wsclient = wsclient
	stream.HandleUnexpected(CreateBytesHandler(func(event *depth.PartialDepthEvent, err error) {
		if err != nil {
			handler(nil, err)
			return
		}

		handler(event.PartialDepthData(), nil)
	}))
	return stream
}

func (s *Stream) HandleBookDepth5Combined(streamName string, handler WsPartialDepthHandler) {
	s.RegisterEventHandler(
		streamName,
		CreateBytesHandler(func(event *depth.PartialDepthEvent, err error) {
			if err != nil {
				handler(nil, err)
				return
			}

			handler(event.PartialDepthData(), nil)
		}),
	)
}
