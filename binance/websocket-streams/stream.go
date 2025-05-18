package websocketstreams

import (
	"context"
	"fmt"
	"goex/binance"
	"goex/binance/models"
	"goex/binance/variables"
	"strings"

	"github.com/CrazyThursdayV50/pkgo/builtin/collector"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
)

type Streamer interface {
	StreamName() string
}

func newStream(streamer Streamer, ctx context.Context, logger log.Logger, handler client.MessageHandler) *binance.Client {
	return binance.NewClient(ctx, logger, fmt.Sprintf(variables.StreamURL(), streamer.StreamName()), handler)
}

func newCombinedStream(streamers []Streamer, ctx context.Context, logger log.Logger, handler client.MessageHandler) *binance.Client {
	return binance.NewClient(ctx, logger, fmt.Sprintf(variables.CombinedStreamURL(), strings.Join(collector.Slice(streamers, func(_ int, streamer Streamer) (bool, string) { return true, streamer.StreamName() }), "/")), handler)
}

type partialBookDepthStreamer struct {
	symbol string
	level  int
}

func (s partialBookDepthStreamer) StreamName() string {
	return fmt.Sprintf("%s@depth%d", strings.ToLower(s.symbol), s.level)

}

type partialBookDepth100msStreamer struct {
	symbol string
	level  int
}

func (s partialBookDepth100msStreamer) StreamName() string {
	return fmt.Sprintf("%s@depth%d@100ms", strings.ToLower(s.symbol), s.level)
}

func PartialBookDepth5Stream(ctx context.Context, logger log.Logger, symbol string, handler WsPartialDepthHandler) *binance.Client {
	client := newStream(
		partialBookDepth100msStreamer{symbol: symbol, level: 5},
		ctx,
		logger,
		func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			var event models.PartialDepthEvent
			err := json.JSON().Unmarshal(b, &event)
			if err != nil {
				f(err)
				return client.BinaryMessage, nil
			}

			handler(event.PartialDepthData())
			return client.BinaryMessage, nil
		})

	client.Run()
	return client
}

func PartialBookDepth5CombinedStream(ctx context.Context, logger log.Logger, symbols []string, handler WsPartialDepthCombinedHandler) *binance.Client {
	client := newCombinedStream(
		collector.Slice(symbols, func(_ int, symbol string) (bool, Streamer) {
			return true, partialBookDepth100msStreamer{symbol: symbol, level: 5}
		}),
		ctx,
		logger,
		func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			// l.Infof("data: %s", b)
			var event models.PartialDepthCombinedEvent
			err := json.JSON().Unmarshal(b, &event)
			if err != nil {
				f(err)
				return client.BinaryMessage, nil
			}

			handler(event.PartialDepthCombinedData())
			return client.BinaryMessage, nil
		})

	client.Run()
	return client
}
