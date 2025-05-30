package stream

import (
	"context"
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/client"
	"github.com/CrazyThursdayV50/goex/binance/websocket-streams/models/klines"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
)

const eventNameKlines = "%s@kline_%s"
const eventNameKlinesTimeZone = "%s@kline_%s_%s"

type klinesStreamer struct {
	symbol   string
	interval string
	timezone string
	logger   log.Logger
	handler  WsHandlerKlines
}

func (s *klinesStreamer) Symbol(value string) *klinesStreamer {
	s.symbol = value
	return s
}

func (s *klinesStreamer) Interval(value string) *klinesStreamer {
	s.interval = value
	return s
}

func (s *klinesStreamer) TimeZone(value string) *klinesStreamer {
	s.timezone = value
	return s
}

func (s klinesStreamer) StreamName() string {
	if s.timezone == "" {
		return fmt.Sprintf(eventNameKlines, strings.ToLower(s.symbol), s.interval)
	}

	return fmt.Sprintf(eventNameKlinesTimeZone, strings.ToLower(s.symbol), s.interval, s.timezone)
}

func KlinesStream(logger log.Logger, handler WsHandlerKlines) *klinesStreamer {
	return &klinesStreamer{
		logger:  logger,
		handler: handler,
	}
}

func (s *klinesStreamer) Connect(ctx context.Context) *client.Client {
	client := newStream(
		s,
		s.logger,
		func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			var event klines.KlinesResult
			err := json.JSON().Unmarshal(b, &event)
			if err != nil {
				f(err)
				return BinaryMessage, nil
			}

			s.handler(&event.Klines)
			return BinaryMessage, nil
		})

	return client
}

// KlinesStream K线数据流
func (ws *WebSocketStreams) KlinesStream(logger log.Logger, handler WsHandlerKlines) *klinesStreamer {
	return KlinesStream(logger, handler)
}
