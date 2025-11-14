package binance

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives"
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

func (Binance) Derivatives() derivatives.DerivativesEntry {
	return derivatives.DerivativesEntry{}
}
