// Package klines provides models and parameters for Binance WebSocket API Klines endpoint.
package klines

import (
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models"
	"github.com/CrazyThursdayV50/pkgo/json"
)

type ParamsData struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime uint64 `json:"startTime,omitempty"`
	EndTime   uint64 `json:"endTime,omitempty"`
	TimeZone  string `json:"timeZone,omitempty"` // [-12:00 ~ +14:00]
	Limit     int    `json:"limit,omitempty"`
}

func defaultParamsData() *ParamsData {
	return &ParamsData{
		Symbol:   "BTCUSDT",
		Interval: Min1.String(),
		Limit:    500,
	}
}

type Params = models.WsAPIParams[*ParamsData]

func NewParams() *Params {
	return &Params{
		Method: "klines",
		Params: defaultParamsData(),
	}
}

/*
[
  1655971200000,      // 这根K线的起始时间
  "0.01086000",       // 这根K线期间第一笔成交价
  "0.01086600",       // 这根K线期间最高成交价
  "0.01083600",       // 这根K线期间最低成交价
  "0.01083800",       // 这根K线期间末一笔成交价
  "2290.53800000",    // 这根K线期间成交量
  1655974799999,      // 这根K线的结束时间
  "24.85074442",      // 这根K线期间成交额
  2283,               // 这根K线期间成交笔数
  "1171.64000000",    // 主动买入的成交量
  "12.71225884",      // 主动买入的成交额
  "0"                 // 忽略此参数
]
*/

type originResultData []any

func (d originResultData) OpenTs() int64 {
	return int64(d[0].(float64))
}

func (d originResultData) Open() string {
	return d[1].(string)
}

func (d originResultData) High() string {
	return d[2].(string)
}

func (d originResultData) Low() string {
	return d[3].(string)
}

func (d originResultData) Close() string {
	return d[4].(string)
}

func (d originResultData) Volume() string {
	return d[5].(string)
}

func (d originResultData) CloseTs() int64 {
	return int64(d[6].(float64))
}

func (d originResultData) Amount() string {
	return d[7].(string)
}

func (d originResultData) TradeCount() int64 {
	return int64(d[8].(float64))
}

func (d originResultData) VolumeBuy() string {
	return d[9].(string)
}

func (d originResultData) AmountBuy() string {
	return d[10].(string)
}

func (d originResultData) Ignore() string {
	return d[11].(string)
}

type Kline struct {
	OpenTs     int64  `json:"openTs"`
	CloseTs    int64  `json:"closeTs"`
	Open       string `json:"open"`
	Close      string `json:"close"`
	High       string `json:"high"`
	Low        string `json:"low"`
	Volume     string `json:"volume"`
	Amount     string `json:"amount"`
	TradeCount int64  `json:"tradeCount"`
	VolumeBuy  string `json:"volumeBuy"`
	AmountBuy  string `json:"amountBuy"`
}

func (d *Kline) UnmarshalJSON(data []byte) error {
	var origin originResultData
	err := json.JSON().Unmarshal(data, &origin)
	if err != nil {
		return err
	}

	d.OpenTs = origin.OpenTs()
	d.CloseTs = origin.CloseTs()
	d.Open = origin.Open()
	d.Close = origin.Close()
	d.High = origin.High()
	d.Low = origin.Low()
	d.Volume = origin.Volume()
	d.Amount = origin.Amount()
	d.TradeCount = origin.TradeCount()
	d.VolumeBuy = origin.VolumeBuy()
	d.AmountBuy = origin.AmountBuy()
	return nil
}

type ResultData []Kline
