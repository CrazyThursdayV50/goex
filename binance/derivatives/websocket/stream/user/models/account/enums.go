package account

// EventReasonType 定义事件原因类型的自定义类型
type EventReasonType string

// 定义 EventReasonType 的所有常量值，对应 JSON 字段 "m" 的所有可能类型。
const (
	EventReasonUnknown EventReasonType = "UNKNOWN"

	EventReasonDeposit             EventReasonType = "DEPOSIT"
	EventReasonWithdraw            EventReasonType = "WITHDRAW"
	EventReasonOrder               EventReasonType = "ORDER"
	EventReasonFundingFee          EventReasonType = "FUNDING_FEE"
	EventReasonWithdrawReject      EventReasonType = "WITHDRAW_REJECT"
	EventReasonAdjustment          EventReasonType = "ADJUSTMENT"
	EventReasonInsuranceClear      EventReasonType = "INSURANCE_CLEAR"
	EventReasonAdminDeposit        EventReasonType = "ADMIN_DEPOSIT"
	EventReasonAdminWithdraw       EventReasonType = "ADMIN_DEPOSIT"
	EventReasonMarginTransfer      EventReasonType = "MARGIN_TRANSFER"
	EventReasonMarginTypeChange    EventReasonType = "MARGIN_TYPE_CHANGE"
	EventReasonAssetTransfer       EventReasonType = "ASSET_TRANSFER"
	EventReasonOptionsPremiumFee   EventReasonType = "OPTIONS_PREMIUM_FEE"
	EventReasonOptionsSettleProfit EventReasonType = "OPTIONS_SETTLE_PROFIT"
	EventReasonAutoExchange        EventReasonType = "AUTO_EXCHANGE"
	EventReasonCoinSwapDeposit     EventReasonType = "COIN_SWAP_DEPOSIT"
	EventReasonCoinSwapWithdraw    EventReasonType = "COIN_SWAP_WITHDRAW"
)
