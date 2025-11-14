package klines

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models"
)

const (
	Event            = "kline"
	streamNameFormat = "%s@kline_%s"
)

func StreamName(symbol, interval string) string {
	return fmt.Sprintf(streamNameFormat, strings.ToLower(symbol), interval)
}

type Result struct {
	models.BaseResult
	Data KlineData `json:"k"`
}

type KlineData struct {
	OpenTime       int64  `json:"t"`
	CloseTime      int64  `json:"T"`
	Symbol         string `json:"s"`
	Interval       string `json:"i"`
	FirstTradeID   int64  `json:"f"`
	LastTradeID    int64  `json:"L"`
	Open           string `json:"o"`
	Close          string `json:"c"`
	High           string `json:"h"`
	Low            string `json:"l"`
	Volume         string `json:"v"`
	Amount         string `json:"q"`
	TradeCount     int64  `json:"n"`
	IsFinal        bool   `json:"x"`
	TakerBuyVolume string `json:"V"`
	TakerBuyAmount string `json:"Q"`
}
