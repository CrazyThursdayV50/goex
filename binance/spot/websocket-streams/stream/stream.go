package stream

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/spot"
	"github.com/CrazyThursdayV50/goex/infra/wsclient"
	"github.com/CrazyThursdayV50/pkgo/builtin/collector"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type Streamer interface {
	StreamName() string
}

func newStream(streamer Streamer, logger log.Logger, handler MessageHandler) *wsclient.Client {
	return wsclient.NewClient(
		logger,
		fmt.Sprintf(spot.WsStream().Endpoint(), streamer.StreamName()),
		variables.GetProxy(),
		handler,
		nil,
	)
}

func newCombinedStream(logger log.Logger, streamers []Streamer, handler MessageHandler) *wsclient.Client {
	return wsclient.NewClient(
		logger,
		fmt.Sprintf(spot.WsStream().EndpointCombined(),
			strings.Join(collector.Slice(streamers, func(_ int, streamer Streamer) (bool, string) { return true, streamer.StreamName() }), "/"),
		),
		variables.GetProxy(),
		handler,
		nil,
	)
}

// NewWebSocketStreams 创建 WebSocket 流客户端的便捷方法
type WebSocketStreams struct{}
