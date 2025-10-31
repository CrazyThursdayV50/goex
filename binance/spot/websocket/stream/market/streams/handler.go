package streams

import (
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models/bookticker"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models/depth"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models/kline"
)

type WsPartialDepthHandler = func(*depth.PartialDepthData, error)

// type WsPartialDepthCombinedHandler = func(*depth.PartialDepthCombinedData, error)
type WsIndividualSymbolBookTickerHandler = func(*bookticker.IndividualSymbolBookTicker, error)
type WsHandlerKlines = func(*kline.Data, error)
