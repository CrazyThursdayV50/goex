package account

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"

type StatusData struct{}

func Status() *models.Request {
	return &models.Request{
		Method: "v2/account.status",
	}
}

// StatusResultData 表示账户总览信息（通常仅针对USDT资产）
type StatusResultData struct {
	TotalInitialMargin          string            `json:"totalInitialMargin"`          // 当前所需起始保证金总额 (存在逐仓请忽略), 仅计算USDT资产
	TotalMaintMargin            string            `json:"totalMaintMargin"`            // 维持保证金总额, 仅计算USDT资产
	TotalWalletBalance          string            `json:"totalWalletBalance"`          // 账户总余额, 仅计算USDT资产
	TotalUnrealizedProfit       string            `json:"totalUnrealizedProfit"`       // 持仓未实现盈亏总额, 仅计算USDT资产
	TotalMarginBalance          string            `json:"totalMarginBalance"`          // 保证金总余额, 仅计算USDT资产
	TotalPositionInitialMargin  string            `json:"totalPositionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格), 仅计算USDT资产
	TotalOpenOrderInitialMargin string            `json:"totalOpenOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格), 仅计算USDT资产
	TotalCrossWalletBalance     string            `json:"totalCrossWalletBalance"`     // 全仓账户余额, 仅计算USDT资产
	TotalCrossUnPnl             string            `json:"totalCrossUnPnl"`             // 全仓持仓未实现盈亏总额, 仅计算USDT资产
	AvailableBalance            string            `json:"availableBalance"`            // 可用余额, 仅计算USDT资产
	MaxWithdrawAmount           string            `json:"maxWithdrawAmount"`           // 最大可转出余额, 仅计算USDT资产
	Assets                      []AccountAssets   `json:"assets"`
	Positions                   []AccountPosition `json:"positions"`
}

// AccountAssets 表示账户资产信息
type AccountAssets struct {
	Asset                  string `json:"asset"`                  // 资产
	WalletBalance          string `json:"walletBalance"`          // 余额
	UnrealizedProfit       string `json:"unrealizedProfit"`       // 未实现盈亏
	MarginBalance          string `json:"marginBalance"`          // 保证金余额
	MaintMargin            string `json:"maintMargin"`            // 维持保证金
	InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金
	PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
	CrossWalletBalance     string `json:"crossWalletBalance"`     // 全仓账户余额
	CrossUnPnl             string `json:"crossUnPnl"`             // 全仓持仓未实现盈亏
	AvailableBalance       string `json:"availableBalance"`       // 可用余额
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`      // 最大可转出余额
	UpdateTime             int64  `json:"updateTime"`             // 更新时间
}

// AccountPosition 表示简化的持仓信息
type AccountPosition struct {
	Symbol           string `json:"symbol"`           // 交易对
	PositionSide     string `json:"positionSide"`     // 持仓方向
	PositionAmt      string `json:"positionAmt"`      // 持仓数量
	UnrealizedProfit string `json:"unrealizedProfit"` // 持仓未实现盈亏
	IsolatedMargin   string `json:"isolatedMargin"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
	InitialMargin    string `json:"initialMargin"` // 持仓所需起始保证金(基于最新标记价格)
	MaintMargin      string `json:"maintMargin"`   // 维持保证金 / 当前杠杆下用户可用的最大名义价值
	UpdateTime       int64  `json:"updateTime"`    // 更新时间
}
