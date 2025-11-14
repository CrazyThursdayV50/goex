package trade

import userdata "github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/useData"

type MarginTypeParams struct {
	Symbol     string              `json:"symbol"`
	MarginType userdata.MarginType `json:"marginType"`
	RecvWindow int64               `json:"recvWindow,omitempty"`
}

type MarginTypeResult struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
