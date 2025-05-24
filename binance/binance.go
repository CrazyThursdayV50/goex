package binance

import (
	"context"

	websocketapi "github.com/CrazyThursdayV50/goex/binance/websocket-api"
	websocketstreams "github.com/CrazyThursdayV50/goex/binance/websocket-streams"
	"github.com/CrazyThursdayV50/pkgo/log"
)

// NewWebSocketAPI 创建一个新的 WebSocket API 客户端
func NewWebSocketAPI(ctx context.Context, logger log.Logger, apiKey, secretKey string) *websocketapi.API {
	return websocketapi.New(ctx, logger, apiKey, secretKey)
}

// NewWebSocketStreams 创建 WebSocket 流客户端的便捷方法
type WebSocketStreams struct{}

func NewWebSocketStreams() *WebSocketStreams {
	return &WebSocketStreams{}
}

// PartialBookDepth5Stream 创建5档深度数据流
func (ws *WebSocketStreams) PartialBookDepth5Stream(ctx context.Context, logger log.Logger, symbol string, handler websocketstreams.WsPartialDepthHandler) *websocketstreams.Client {
	return websocketstreams.PartialBookDepth5Stream(ctx, logger, symbol, handler)
}

// PartialBookDepth5CombinedStream 创建组合5档深度数据流
func (ws *WebSocketStreams) PartialBookDepth5CombinedStream(ctx context.Context, logger log.Logger, symbols []string, handler websocketstreams.WsPartialDepthCombinedHandler) *websocketstreams.Client {
	return websocketstreams.PartialBookDepth5CombinedStream(ctx, logger, symbols, handler)
}

// IndividualSymbolBookTickerStream 创建个股最佳买卖价流
func (ws *WebSocketStreams) IndividualSymbolBookTickerStream(ctx context.Context, logger log.Logger, symbols []string, handler websocketstreams.WsIndividualSymbolBookTickerHandler) *websocketstreams.Client {
	return websocketstreams.IndividualSymbolBookTickerStream(ctx, logger, symbols, handler)
} 