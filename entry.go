package goex

import (
	"github.com/CrazyThursdayV50/goex/binance"
)

func Binance() binance.Binance {
	return binance.New()
}
