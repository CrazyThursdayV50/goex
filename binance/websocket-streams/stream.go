package websocketstreams

import (
	"context"
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/websocket-streams/models"
	"github.com/CrazyThursdayV50/pkgo/builtin/collector"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
)

type Streamer interface {
	StreamName() string
}

func newStream(streamer Streamer, ctx context.Context, logger log.Logger, handler client.MessageHandler) *Client {
	return NewClient(ctx, logger, fmt.Sprintf(variables.StreamURL(), streamer.StreamName()), handler)
}

func newCombinedStream(ctx context.Context, logger log.Logger, streamers []Streamer, handler client.MessageHandler) *Client {
	return NewClient(ctx, logger, fmt.Sprintf(variables.CombinedStreamURL(), strings.Join(collector.Slice(streamers, func(_ int, streamer Streamer) (bool, string) { return true, streamer.StreamName() }), "/")), handler)
}

type partialBookDepthStreamer struct {
	symbol string
	level  int
}

func (s partialBookDepthStreamer) StreamName() string {
	return fmt.Sprintf(variables.PARTIAL_BOOK_DEPTH, strings.ToLower(s.symbol), s.level)

}

type partialBookDepth100msStreamer struct {
	symbol string
	level  int
}

func (s partialBookDepth100msStreamer) StreamName() string {
	return fmt.Sprintf(variables.PARTIAL_BOOK_DEPTH_100ms, strings.ToLower(s.symbol), s.level)
}

func PartialBookDepth5Stream(ctx context.Context, logger log.Logger, symbol string, handler WsPartialDepthHandler) *Client {
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

	err := client.Run()
	if err != nil {
		panic(err)
	}
	return client
}

func PartialBookDepth5CombinedStream(ctx context.Context, logger log.Logger, symbols []string, handler WsPartialDepthCombinedHandler) *Client {
	client := newCombinedStream(
		ctx,
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
				return client.BinaryMessage, nil
			}

			handler(event.PartialDepthCombinedData())
			return client.BinaryMessage, nil
		})

	err := client.Run()
	if err != nil {
		panic(err)
	}

	return client
}
