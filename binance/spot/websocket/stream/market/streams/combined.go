package streams

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/spot"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
)

func (s *Stream) Combined(streamNames []string) *Stream {
	stream := s.Clone()
	wsclient := client.NewClient(
		stream.logger,
		fmt.Sprintf(spot.WsStream().EndpointCombined(), strings.Join(streamNames, "/")),
		variables.GetProxy(),
		stream.handler,
		nil,
	)
	stream.wsclient = wsclient
	return stream
}
