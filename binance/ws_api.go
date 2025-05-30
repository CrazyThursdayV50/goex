package binance

import (
	"github.com/CrazyThursdayV50/goex/binance/websocket-api/api"
	"github.com/CrazyThursdayV50/pkgo/log"
)

// NewWebSocketAPI 创建一个新的 WebSocket API 客户端
type WebsocketAPI = api.API

func NewWebSocketAPI(logger log.Logger, apiKey, secretKey string) *WebsocketAPI {
	return api.New(logger, apiKey, secretKey)
}
