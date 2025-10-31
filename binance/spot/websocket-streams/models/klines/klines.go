package klines

import "github.com/CrazyThursdayV50/goex/binance/spot/websocket-streams/models"

/*
{
  "e": "kline",          // 事件类型
  "E": 1672515782136,    // 事件时间
  "s": "BNBBTC",         // 交易对
  "k": {
    "t": 1672515780000,  // 这根K线的起始时间
    "T": 1672515839999,  // 这根K线的结束时间
    "s": "BNBBTC",       // 交易对
    "i": "1m",           // K线间隔
    "f": 100,            // 这根K线期间第一笔成交ID
    "L": 200,            // 这根K线期间末一笔成交ID
    "o": "0.0010",       // 这根K线期间第一笔成交价
    "c": "0.0020",       // 这根K线期间末一笔成交价
    "h": "0.0025",       // 这根K线期间最高成交价
    "l": "0.0015",       // 这根K线期间最低成交价
    "v": "1000",         // 这根K线期间成交量
    "n": 100,            // 这根K线期间成交数量
    "x": false,          // 这根K线是否完结（是否已经开始下一根K线）
    "q": "1.0000",       // 这根K线期间成交额
    "V": "500",          // 主动买入的成交量
    "Q": "0.500",        // 主动买入的成交额
    "B": "123456"        // 忽略此参数
  }
}
*/

type Data struct {
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

type KlinesResult struct {
	models.StreamBase
	Klines Data `json:"k"`
}
