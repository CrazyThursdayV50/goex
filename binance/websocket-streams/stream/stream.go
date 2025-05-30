package stream

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/client"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/pkgo/builtin/collector"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type Streamer interface {
	StreamName() string
}

func newStream(streamer Streamer, logger log.Logger, handler MessageHandler) *client.Client {
	return client.NewClient(logger, fmt.Sprintf(variables.StreamURL(), streamer.StreamName()), handler, nil)
}

func newCombinedStream(logger log.Logger, streamers []Streamer, handler MessageHandler) *client.Client {
	return client.NewClient(
		logger,
		fmt.Sprintf(variables.CombinedStreamURL(),
			strings.Join(collector.Slice(streamers, func(_ int, streamer Streamer) (bool, string) { return true, streamer.StreamName() }), "/"),
		),
		handler,
		nil,
	)
}

// NewWebSocketStreams 创建 WebSocket 流客户端的便捷方法
type WebSocketStreams struct{}
