package stream

import (
	"context"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-streams/models"
	"github.com/CrazyThursdayV50/goex/binance/variables/spot"
	"github.com/CrazyThursdayV50/goex/infra/wsclient"
	"github.com/CrazyThursdayV50/pkgo/builtin/collector"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type partialBookDepthStreamer struct {
	symbol string
	level  int
}

func (s partialBookDepthStreamer) StreamName() string {
	return spot.WsStream().PartialBookDepth(strings.ToLower(s.symbol), s.level)

}

type partialBookDepth100msStreamer struct {
	symbol string
	level  int
}

func (s partialBookDepth100msStreamer) StreamName() string {
	return spot.WsStream().PartialBookDepth100ms(strings.ToLower(s.symbol), s.level)
}

func PartialBookDepth5Stream(ctx context.Context, logger log.Logger, symbol string, handler WsPartialDepthHandler) *wsclient.Client {
	client := newStream(
		partialBookDepth100msStreamer{symbol: symbol, level: 5},
		logger,
		func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			var event models.PartialDepthEvent
			err := json.JSON().Unmarshal(b, &event)
			if err != nil {
				f(err)
				return BinaryMessage, nil
			}

			handler(event.PartialDepthData())
			return BinaryMessage, nil
		})

	return client
}

func PartialBookDepth5CombinedStream(ctx context.Context, logger log.Logger, symbols []string, handler WsPartialDepthCombinedHandler) *wsclient.Client {
	client := newCombinedStream(
		logger,
		collector.Slice(symbols, func(_ int, symbol string) (bool, Streamer) {
			return true, partialBookDepth100msStreamer{symbol: symbol, level: 5}
		}),
		func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			// l.Infof("data: %s", b)
			var event models.PartialDepthCombinedEvent
			err := json.JSON().Unmarshal(b, &event)
			if err != nil {
				f(err)
				return BinaryMessage, nil
			}

			handler(event.PartialDepthCombinedData())
			return BinaryMessage, nil
		})

	return client
}

// PartialBookDepth5Stream 创建5档深度数据流
func (ws *WebSocketStreams) PartialBookDepth5Stream(ctx context.Context, logger log.Logger, symbol string, handler WsPartialDepthHandler) *wsclient.Client {
	return PartialBookDepth5Stream(ctx, logger, symbol, handler)
}

// PartialBookDepth5CombinedStream 创建组合5档深度数据流
func (ws *WebSocketStreams) PartialBookDepth5CombinedStream(ctx context.Context, logger log.Logger, symbols []string, handler WsPartialDepthCombinedHandler) *wsclient.Client {
	return PartialBookDepth5CombinedStream(ctx, logger, symbols, handler)
}
