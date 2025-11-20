package account

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models"

const (
	EventName = "ACCOUNT_UPDATE"
)

/*
 当账户信息有变动时，会推送此事件：
	仅当账户信息有变动时(包括资金、仓位、保证金模式等发生变化)，才会推送此事件；
	订单状态变化没有引起账户和持仓变化的，不会推送此事件；
	position 信息：仅当 symbol 仓位有变动时推送。
*/

/*
"FUNDING FEE" 引起的资金余额变化，仅推送简略事件：

当用户某全仓持仓发生"FUNDING FEE"时，事件ACCOUNT_UPDATE将只会推送相关的用户资产余额信息B(仅推送 FUNDING FEE 发生相关的资产余额信息)，而不会推送任何持仓信息P。
当用户某逐仓仓持仓发生"FUNDING FEE"时，事件ACCOUNT_UPDATE将只会推送相关的用户资产余额信息B(仅推送"FUNDING FEE"所使用的资产余额信息)，和相关的持仓信息P(仅推送这笔"FUNDING FEE"发生所在的持仓信息)，其余持仓信息不会被推送。
*/

type Event struct {
	models.BaseEvent
	Account EventData `json:"a"`
}

// EventData 账户更新事件数据
type EventData struct {
	// EventReason 对应 JSON 中的 "m" 键, e.g., "ORDER" (事件推出原因)
	EventReason EventReasonType `json:"m"`
	// BalanceInformationList 对应 JSON 中的 "B" 键 (余额信息列表)
	BalanceInformationList []BalanceInfo `json:"B"`
	// PositionInformationList 对应 JSON 中的 "P" 键 (仓位信息列表)
	PositionInformationList []PositionInfo `json:"P"`
}

// BalanceInfo 账户余额信息
type BalanceInfo struct {
	// AssetName 对应 JSON 中的 "a" 键, e.g., "USDT" (资产名称)
	AssetName string `json:"a"`
	// WalletBalance 对应 JSON 中的 "wb" 键 (钱包余额，使用 string 保持精度)
	WalletBalance string `json:"wb"`
	// CrossWalletBalance 对应 JSON 中的 "cw" 键 (除去逐仓仓位保证金的钱包余额，使用 string)
	CrossWalletBalance string `json:"cw"`
	// BalanceChange 对应 JSON 中的 "bc" 键 (除去盈亏与交易手续费以外的钱包余额改变量，使用 string)
	// 字段"bc"代表了钱包余额的改变量，即 balance change，但注意其不包含仓位盈亏及交易手续费。
	BalanceChange string `json:"bc"`
}

// PositionInfo 仓位信息
type PositionInfo struct {
	// Symbol 对应 JSON 中的 "s" 键, e.g., "BTCUSDT" (交易对)
	Symbol string `json:"s"`
	// PositionAmount 对应 JSON 中的 "pa" 键 (仓位/持仓量，使用 string)
	PositionAmount string `json:"pa"`
	// EntryPrice 对应 JSON 中的 "ep" 键 (入仓价格，使用 string)
	EntryPrice string `json:"ep"`
	// BreakEvenPrice 对应 JSON 中的 "bep" 键 (盈亏平衡价，使用 string)
	BreakEvenPrice string `json:"bep"`
	// CumulativeRealizedProfitLoss 对应 JSON 中的 "cr" 键 ((费前)累计实现损益，使用 string)
	CumulativeRealizedProfitLoss string `json:"cr"`
	// UnrealizedProfitLoss 对应 JSON 中的 "up" 键 (持仓未实现盈亏，使用 string)
	UnrealizedProfitLoss string `json:"up"`
	// MarginType 对应 JSON 中的 "mt" 键 (保证金模式, e.g., "isolated")
	MarginType string `json:"mt"`
	// IsolatedWalletMargin 对应 JSON 中的 "iw" 键 (若为逐仓，仓位保证金，使用 string)
	IsolatedWalletMargin string `json:"iw"`
	// PositionSide 对应 JSON 中的 "ps" 键 (持仓方向, e.g., "BOTH", "LONG", "SHORT")
	PositionSide string `json:"ps"`
}
