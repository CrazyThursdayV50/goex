package binance

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api"
	"github.com/CrazyThursdayV50/goex/binance/spot"
)

type Binance struct {
}

func New() Binance {
	return Binance{}
}

func (Binance) Spot() spot.SpotEntry {
	return spot.New()
}

func (Binance) Derivatives() api.API {
	return api.New()
}
