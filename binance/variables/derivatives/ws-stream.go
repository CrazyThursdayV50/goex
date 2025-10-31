package derivatives

import (
	"github.com/CrazyThursdayV50/goex/binance/variables"
)

const (
	urlBaseWebSocketStream      = "wss://fstream.binance.com"
	urlBaseWebSocketStreamTest  = "wss://demo-fstream.binance.com"
	pathWebSocketStream         = "/ws"
	pathWebSocketCombinedStream = "/stream"
)

type stream struct{}

func WsStream() stream { return stream{} }

func streamBaseURL() string {
	if variables.IsTest() {
		return urlBaseWebSocketStreamTest
	}
	return urlBaseWebSocketStream
}

func (stream) Endpoint() string         { return streamBaseURL() + pathWebSocketStream }
func (stream) EndpointCombined() string { return streamBaseURL() + pathWebSocketCombinedStream }
