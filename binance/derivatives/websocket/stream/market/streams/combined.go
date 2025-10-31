package streams

import (
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
)

func (s *Stream) Combined() *Stream {
	stream := s.Clone()
	stream.wsclient = client.NewClient(
		stream.logger,
		derivatives.WsStream().EndpointCombined(),
		variables.GetProxy(),
		stream.handler,
		nil,
	)

	return stream
}
