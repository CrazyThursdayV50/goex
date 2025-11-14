package position

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"

type RequestData struct {
	Symbol string `json:"symbol"`
}

func Params() *models.Request {
	return &models.Request{
		Method: "v2/account.position",
	}
}

/*
	 {
		    "symbol": "ADAUSDT",
		    "positionSide": "BOTH",               // 持仓方向
		    "positionAmt": "30",
		    "entryPrice": "0.385",
		    "breakEvenPrice": "0.385077",
		    "markPrice": "0.41047590",
		    "unRealizedProfit": "0.76427700",     // 持仓未实现盈亏
		    "liquidationPrice": "0",
		    "isolatedMargin": "0",
		    "notional": "12.31427700",
		    "marginAsset": "USDT",
		    "isolatedWallet": "0",
		    "initialMargin": "0.61571385",        // 初始保证金
		    "maintMargin": "0.08004280",          // 维持保证金
		    "positionInitialMargin": "0.61571385",// 仓位初始保证金
		    "openOrderInitialMargin": "0",        // 订单初始保证金
		    "adl": 2,
		    "bidNotional": "0",
		    "askNotional": "0",
		    "updateTime": 1720736417660           // 更新时间
		}
*/
type Position struct {
	Symbol                 string `json:"symbol"`
	PositionSide           string `json:"positionSide"`           // 持仓方向
	PositionAmt            string `json:"positionAmt"`            // 仓位数量
	EntryPrice             string `json:"entryPrice"`             // 开仓均价
	BreakEvenPrice         string `json:"breakEvenPrice"`         // 损益两平价
	MarkPrice              string `json:"markPrice"`              // 标记价格
	UnRealizedProfit       string `json:"unRealizedProfit"`       // 持仓未实现盈亏
	LiquidationPrice       string `json:"liquidationPrice"`       // 强平价格
	IsolatedMargin         string `json:"isolatedMargin"`         // 逐仓模式仓位总保证金(即钱包余额与未实现盈亏之和)
	Notional               string `json:"notional"`               // 仓位名义价值
	MarginAsset            string `json:"marginAsset"`            // 保证金资产币种
	IsolatedWallet         string `json:"isolatedWallet"`         // 逐仓模式仓位钱包余额
	InitialMargin          string `json:"initialMargin"`          // 初始保证金
	MaintMargin            string `json:"maintMargin"`            // 维持保证金
	PositionInitialMargin  string `json:"positionInitialMargin"`  // 仓位初始保证金
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 订单初始保证金
	Adl                    int    `json:"adl"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	UpdateTime             int64  `json:"updateTime"` // 更新时间 (通常使用 int64 来存储时间戳)
}

type ResultData []Position
