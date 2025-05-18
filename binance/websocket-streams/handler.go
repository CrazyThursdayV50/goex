package websocketstreams

import "github.com/CrazyThursdayV50/goex/binance/models"

type WsPartialDepthHandler = func(*models.PartialDepthData)
type WsPartialDepthCombinedHandler = func(*models.PartialDepthCombinedData)
