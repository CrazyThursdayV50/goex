package goex

import (
	"github.com/CrazyThursdayV50/goex/binance"
)

type binanceEntry struct{}

func (binanceEntry) Spot() binance.SpotEntry {
	return binance.Spot()
}

func Binance() binanceEntry {
	return binanceEntry{}
}
