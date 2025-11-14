package trade

type OpenLeverageParams struct {
	Symbol     string `json:"symbol"`
	Leverage   int64  `json:"leverage"` // 目标杠杆倍数：1 到 125 整数
	RecvWindow int64  `json:"recvWindow,omitempty"`
}

type OpenLeverageResult struct {
	// 杠杆倍数
	Leverage int64 `json:"leverage"`

	// 当前杠杆倍数下允许的最大名义价值
	MaxNotionalValue string `json:"maxNotionalValue"`

	// 交易对
	Symbol string `json:"symbol"`
}
