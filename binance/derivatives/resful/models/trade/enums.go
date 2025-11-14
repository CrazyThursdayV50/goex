package trade

import userdata "github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/useData"

type MarginType = userdata.MarginType

const (
	MarginTypeCross    MarginType = "CROSSED"
	MarginTypeIsolated MarginType = "ISOLATED"
)
