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

type bookTickerStreamer struct {
	symbol string
}

func (s bookTickerStreamer) StreamName() string {
	return spot.WsStream().IndividualBookTicker(strings.ToLower(s.symbol))
}

// 最佳买卖数据
func IndividualSymbolBookTickerStream(logger log.Logger, symbols []string, handler WsIndividualSymbolBookTickerHandler) *wsclient.Client {
	client := newCombinedStream(
		logger,
		collector.Slice(symbols, func(_ int, symbol string) (bool, Streamer) {
			return true, bookTickerStreamer{symbol: symbol}
		}),
		func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			var event models.IndividualSymbolBookTickerEvent
			err := json.JSON().Unmarshal(b, &event)
			if err != nil {
				f(err)
				return BinaryMessage, nil
			}

			handler(event.Data)
			return BinaryMessage, nil
		})

	return client
}

// IndividualSymbolBookTickerStream 创建个股最佳买卖价流
func (ws *WebSocketStreams) IndividualSymbolBookTickerStream(logger log.Logger, symbols []string, handler WsIndividualSymbolBookTickerHandler) *wsclient.Client {
	return IndividualSymbolBookTickerStream(logger, symbols, handler)
}
