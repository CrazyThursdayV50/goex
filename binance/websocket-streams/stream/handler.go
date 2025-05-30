package stream

import (
	"github.com/CrazyThursdayV50/goex/binance/websocket-streams/models"
	"github.com/CrazyThursdayV50/goex/binance/websocket-streams/models/klines"
)

type WsPartialDepthHandler = func(*models.PartialDepthData)
type WsPartialDepthCombinedHandler = func(*models.PartialDepthCombinedData)
type WsIndividualSymbolBookTickerHandler = func(*models.IndividualSymbolBookTicker)
type WsHandlerKlines = func(*klines.Data)
