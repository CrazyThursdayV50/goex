package goex

import (
	binanceDerivativesRest "github.com/CrazyThursdayV50/goex/binance/derivatives/resful/api"
	binanceDerivativesWebSocket "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/apis"
	binanceDerivativesWebSocketMarket "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/streams"
	binanceSpotWebSocket "github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/apis"
	binanceSpotWebSocketMarket "github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/streams"
)

// binance
// binance spot
type BinanceSpotWebSocketAPI = binanceSpotWebSocket.API
type BinanceSpotWebSocketMarketStream = binanceSpotWebSocketMarket.Stream

// binance derivaties
type BinanceDerivativesWebSocketAPI = binanceDerivativesWebSocket.API
type BinanceDerivativesWebSocketMarketStream = binanceDerivativesWebSocketMarket.Stream
type BinanceDerivativesRestAPI = binanceDerivativesRest.API
