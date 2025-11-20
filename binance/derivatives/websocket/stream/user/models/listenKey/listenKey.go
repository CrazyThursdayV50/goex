package listenkey

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models"

const (
	EventName = "listenKeyExpired"
)

type Event struct {
	models.BaseEvent
	Data
}

type Data struct {
	ListenKey string `json:"listenKey"`
}
