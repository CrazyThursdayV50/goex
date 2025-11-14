package userdata

type SymbolConfigData struct {
	Symbol     string `json:"symbol"`
	RecvWindow int64  `json:"recvWindow,omitempty"`
}

type SymbolConfigResultData struct {
	Symbol           string     `json:"symbol"`
	MarginType       MarginType `json:"marginType"`
	IsAutoAddMargin  bool       `json:"isAutoAddMargin"` // 注意：虽然值是 "false"，但在 JSON 中它是字符串类型
	Leverage         int64      `json:"leverage"`
	MaxNotionalValue string     `json:"maxNotionalValue"`
}

type SymbolConfigResult []SymbolConfigResultData
