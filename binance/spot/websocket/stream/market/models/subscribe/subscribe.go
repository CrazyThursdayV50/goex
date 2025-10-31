package subscribe

import (
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models"
)

type RequestParams []string

func NewSubscribe() *models.Request {
	return &models.Request{
		Method: "SUBSCRIBE",
	}
}

func NewUnsubscribe() *models.Request {
	return &models.Request{
		Method: "UNSUBSCRIBE",
	}
}
