package spot

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/variables"
)

// 本篇所列出的所有wss接口的baseurl为: wss://stream.binance.com:9443 或者 wss://stream.binance.com:443
const urlBaseWebsocketStream = "wss://stream.binance.com:443"
const urlBaseWebsocketStreamTest = "wss://stream.testnet.binance.vision/ws"

// const BASE_URL = "wss://stream.binance.com:9443"
const pathWebsocketStream = "/ws/%s"
const pathWebsocketCombinedStream = "/stream?streams=%s"

// Stream Names: <symbol>@depth<levels> OR <symbol>@depth<levels>@100ms
const partialBookDepth = "%s@depth%d"
const partialBookDepth100ms = "%s@depth%d@100ms"

const individualBookTicker = "%s@bookTicker"
const klines = "%s@kline_%s"
const klinesTimeZone = "%s@kline_%s_%s"

type stream struct{}

func WsStream() stream { return stream{} }

func streamBaseURL() string {
	if variables.IsTest() {
		return urlBaseWebsocketStreamTest
	}

	return urlBaseWebsocketStream
}

func (stream) Endpoint() string {
	return streamBaseURL() + pathWebsocketStream
}

func (stream) EndpointCombined() string {
	return streamBaseURL() + pathWebsocketCombinedStream
}

func (stream) PartialBookDepth(symbol string, level int) string {
	return fmt.Sprintf(partialBookDepth, symbol, level)
}

func (stream) PartialBookDepth100ms(symbol string, level int) string {
	return fmt.Sprintf(partialBookDepth100ms, symbol, level)
}

func (stream) IndividualBookTicker(symbol string) string {
	return fmt.Sprintf(individualBookTicker, symbol)
}

func (stream) Klines(symbol string, interval string) string {
	return fmt.Sprintf(klines, symbol, interval)
}

func (stream) KlinesTimeZone(symbol string, interval string, timezone string) string {
	return fmt.Sprintf(klinesTimeZone, symbol, interval, timezone)
}
