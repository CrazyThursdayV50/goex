package subscribe

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models"
)

type RequestParams []string

type Request = models.Request[RequestParams]

func NewSubscribe() *Request {
	return &Request{
		Method: "SUBSCRIBE",
	}
}

func NewUnsubscribe() *Request {
	return &Request{
		Method: "UNSUBSCRIBE",
	}
}
