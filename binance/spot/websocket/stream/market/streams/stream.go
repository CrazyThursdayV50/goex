package streams

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/spot"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
	"github.com/CrazyThursdayV50/pkgo/builtin/collector"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type Streamer interface {
	StreamName() string
}

func newStream(streamer Streamer, logger log.Logger, handler client.MessageHandler) *client.Client {
	return client.NewClient(
		logger,
		fmt.Sprintf(spot.WsStream().Endpoint(), streamer.StreamName()),
		variables.GetProxy(),
		handler,
		nil,
	)
}

func newCombinedStream(logger log.Logger, streamers []Streamer, handler client.MessageHandler) *client.Client {
	return client.NewClient(
		logger,
		fmt.Sprintf(spot.WsStream().EndpointCombined(),
			strings.Join(collector.Slice(streamers, func(_ int, streamer Streamer) (bool, string) { return true, streamer.StreamName() }), "/"),
		),
		variables.GetProxy(),
		handler,
		nil,
	)
}
