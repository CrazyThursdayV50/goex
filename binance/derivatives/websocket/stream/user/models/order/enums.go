package order

import (
	"encoding/json"
	"strconv"
)

/*
订单方向

BUY 买入
SELL 卖出
订单类型

LIMIT 限价单
MARKET 市价单
STOP 止损限价单
STOP_MARKET 止损市价单
TAKE_PROFIT 止盈限价单
TAKE_PROFIT_MARKET 止盈市价单
TRAILING_STOP_MARKET 跟踪止损单
LIQUIDATION 爆仓
本次事件的具体执行类型

NEW
CANCELED 已撤
CALCULATED 订单 ADL 或爆仓
EXPIRED 订单失效
TRADE 交易
AMENDMENT 订单修改
订单状态

NEW
PARTIALLY_FILLED
FILLED
CANCELED
EXPIRED
EXPIRED_IN_MATCH
有效方式:

GTC
IOC
FOK
GTX
强平和ADL:

若用户因保证金不足发生强平：
c为"autoclose-XXX"，X为"NEW"
若用户保证金充足但被 ADL:
c为“adl_autoclose”，X为“NEW”
过期原因

0: 无，默认值
1: 自成交保护，订单被取消
2: IOC订单无法完全成交，订单被取消
3: IOC订单因自成交保护无法完全成交，订单被取消
4: 只减仓竞争过程中，低优先级的只减仓订单被取消
5: 账户强平过程中，订单被取消
6: 不满足GTE条件，订单被取消
7: Symbol下架，订单被取消
8: 止盈止损单触发后，初始订单被取消
9: 市价订单无法完全成交，订单被取消
*/

// OrderSide 订单方向
type OrderSide string

const (
	OrderSideBuy  = "BUY"
	OrderSideSell = "SELL"
)

// OrderType 定义订单类型
type OrderType string

const (
	OrderTypeLimit              OrderType = "LIMIT"                // 限价单
	OrderTypeMarket             OrderType = "MARKET"               // 市价单
	OrderTypeStopLimit          OrderType = "STOP"                 // 止损限价单
	OrderTypeStopMarket         OrderType = "STOP_MARKET"          // 止损市价单
	OrderTypeTakeProfitLimit    OrderType = "TAKE_PROFIT"          // 止盈限价单
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"   // 止盈市价单
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET" // 跟踪止损单
	OrderTypeLiquidation        OrderType = "LIQUIDATION"          // 爆仓
)

// ExecutionType 定义本次事件的具体执行类型
type ExecutionType string

const (
	ExecutionTypeNew        ExecutionType = "NEW"        // 新订单
	ExecutionTypeCanceled   ExecutionType = "CANCELED"   // 已撤
	ExecutionTypeCalculated ExecutionType = "CALCULATED" // 订单 ADL 或爆仓
	ExecutionTypeExpired    ExecutionType = "EXPIRED"    // 订单失效
	ExecutionTypeTrade      ExecutionType = "TRADE"      // 交易
	ExecutionTypeAmendment  ExecutionType = "AMENDMENT"  // 订单修改
)

// OrderStatus 定义订单状态
type OrderStatus string

const (
	OrderStatusNew             OrderStatus = "NEW"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled          OrderStatus = "FILLED"
	OrderStatusCanceled        OrderStatus = "CANCELED"
	OrderStatusExpired         OrderStatus = "EXPIRED"
	OrderStatusExpiredInMatch  OrderStatus = "EXPIRED_IN_MATCH"
)

// TimeInForceType 定义有效方式
type TimeInForceType string

const (
	TimeInForceGTC TimeInForceType = "GTC" // Good-Till-Cancel，成交为止
	TimeInForceIOC TimeInForceType = "IOC" // Immediate-Or-Cancel，立即成交或取消
	TimeInForceFOK TimeInForceType = "FOK" // Fill-Or-Kill，全部成交或全部取消
	TimeInForceGTX TimeInForceType = "GTX" // Good-Till-Crossing (一般用于止盈止损单，订单被触发前有效)
)

// ExpireReasonCode 定义过期原因（JSON 中可能是整型）
type ExpireReasonCode int

const (
	ExpireReasonNone                      ExpireReasonCode = 0 // 无，默认值
	ExpireReasonSelfTradePrevention       ExpireReasonCode = 1 // 自成交保护，订单被取消
	ExpireReasonIOCNotFullyFilled         ExpireReasonCode = 2 // IOC订单无法完全成交，订单被取消
	ExpireReasonIOCSelfTradePrevention    ExpireReasonCode = 3 // IOC订单因自成交保护无法完全成交，订单被取消
	ExpireReasonReduceOnlyLowPriority     ExpireReasonCode = 4 // 只减仓竞争过程中，低优先级的只减仓订单被取消
	ExpireReasonLiquidationCancel         ExpireReasonCode = 5 // 账户强平过程中，订单被取消
	ExpireReasonGTENotSatisfied           ExpireReasonCode = 6 // 不满足GTE条件，订单被取消
	ExpireReasonSymbolDelisted            ExpireReasonCode = 7 // Symbol下架，订单被取消
	ExpireReasonStopOrTakeProfitTriggered ExpireReasonCode = 8 // 止盈止损单触发后，初始订单被取消
	ExpireReasonMarketOrderNotFullyFilled ExpireReasonCode = 9 // 市价订单无法完全成交，订单被取消
	// 还可以添加一个通用的未知类型
	ExpireReasonUnknown ExpireReasonCode = 999
)

// UnmarshalJSON 实现了 json.Unmarshaler 接口，以便 JSON 整数或字符串能正确解析为 ExpireReasonCode
func (erc *ExpireReasonCode) UnmarshalJSON(data []byte) error {
	var i int
	// 尝试解析为整数
	if err := json.Unmarshal(data, &i); err == nil {
		*erc = ExpireReasonCode(i)
		return nil
	}

	var s string
	// 如果不是整数，尝试解析为字符串（以防 JSON 字段被引号包围）
	if err := json.Unmarshal(data, &s); err == nil {
		if val, err := strconv.Atoi(s); err == nil {
			*erc = ExpireReasonCode(val)
			return nil
		}
	}

	// 如果解析失败，设置为未知
	*erc = ExpireReasonUnknown
	return nil
}

// String 方法用于方便地将代码转换为描述
func (erc ExpireReasonCode) String() string {
	switch erc {
	case ExpireReasonNone:
		return "无 (默认值)"
	case ExpireReasonSelfTradePrevention:
		return "自成交保护，订单被取消"
	case ExpireReasonIOCNotFullyFilled:
		return "IOC订单无法完全成交，订单被取消"
	case ExpireReasonIOCSelfTradePrevention:
		return "IOC订单因自成交保护无法完全成交，订单被取消"
	case ExpireReasonReduceOnlyLowPriority:
		return "只减仓竞争过程中，低优先级的只减仓订单被取消"
	case ExpireReasonLiquidationCancel:
		return "账户强平过程中，订单被取消"
	case ExpireReasonGTENotSatisfied:
		return "不满足GTE条件，订单被取消"
	case ExpireReasonSymbolDelisted:
		return "Symbol下架，订单被取消"
	case ExpireReasonStopOrTakeProfitTriggered:
		return "止盈止损单触发后，初始订单被取消"
	case ExpireReasonMarketOrderNotFullyFilled:
		return "市价订单无法完全成交，订单被取消"
	default:
		return "未知过期原因"
	}
}
