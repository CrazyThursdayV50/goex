package streams

import (
	"fmt"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
	"github.com/gorilla/websocket"
)

func (s *Stream) Combined() *Stream {
	stream := s.Clone()
	stream.wsclient = client.NewClient(
		stream.logger,
		derivatives.WsStream().EndpointCombined(),
		variables.GetProxy(),
		stream.handler,
		func(done <-chan struct{}, conn *websocket.Conn) {
			ticker := time.NewTicker(time.Minute * 15)
			for {
				select {
				case <-done:
					return

				case t := <-ticker.C:
					conn.WriteControl(
						client.PingMessage,
						fmt.Appendf(nil, "%d", t.UnixMilli()),
						time.Now().Add(time.Second*30))

					conn.WriteControl(
						client.PongMessage,
						fmt.Appendf(nil, "%d", t.UnixMilli()),
						time.Now().Add(time.Second*30))
				}
			}
		},
	)

	return stream
}
