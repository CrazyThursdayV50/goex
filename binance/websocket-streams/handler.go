package websocketstreams

import "goex/binance/models"

type WsPartialDepthHandler = func(*models.PartialDepthData)
type WsPartialDepthCombinedHandler = func(*models.PartialDepthCombinedData)
