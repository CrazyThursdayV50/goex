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

type bookTickerStreamer struct {
	symbol string
}

func (s bookTickerStreamer) StreamName() string {
	return fmt.Sprintf(variables.INDIVIDUAL_BOOK_TICKER, strings.ToLower(s.symbol))
}

func IndividualSymbolBookTickerStream(ctx context.Context, logger log.Logger, symbols []string, handler WsIndividualSymbolBookTickerHandler) *Client {
	client := newCombinedStream(
		ctx,
		logger,
		collector.Slice(symbols, func(_ int, symbol string) (bool, Streamer) {
			return true, bookTickerStreamer{symbol: symbol}
		}),
		func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			var event models.IndividualSymbolBookTickerEvent
			err := json.JSON().Unmarshal(b, &event)
			if err != nil {
				f(err)
				return client.BinaryMessage, nil
			}

			handler(event.Data)
			return client.BinaryMessage, nil
		})

	err := client.Run()
	if err != nil {
		panic(err)
	}
	return client
}
