package klines

type Interval string

const (
	Sec1   Interval = "1s"
	Min1   Interval = "1m"
	Min3   Interval = "3m"
	Min5   Interval = "5m"
	Min15  Interval = "15m"
	Min30  Interval = "30m"
	Hour1  Interval = "1h"
	Hour2  Interval = "2h"
	Hour4  Interval = "4h"
	Hour6  Interval = "6h"
	Hour8  Interval = "8h"
	Hour12 Interval = "12h"
	Day1   Interval = "1d"
	Day3   Interval = "3d"
	Week1  Interval = "1w"
	Month1 Interval = "1M"
)

func (i Interval) String() string {
	return string(i)
}
