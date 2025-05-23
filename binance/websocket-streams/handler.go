package websocketstreams

import "github.com/CrazyThursdayV50/goex/binance/websocket-streams/models"

type WsPartialDepthHandler = func(*models.PartialDepthData)
type WsPartialDepthCombinedHandler = func(*models.PartialDepthCombinedData)
type WsIndividualSymbolBookTickerHandler = func(*models.IndividualSymbolBookTicker)
