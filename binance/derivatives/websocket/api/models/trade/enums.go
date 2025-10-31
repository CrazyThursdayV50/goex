package trade

type OrderSide string

const (
	SIDE_BUY  OrderSide = "BUY"
	SIDE_SELL OrderSide = "SELL"
)

type PositionSide string

const (
	// 单向持仓模式下
	POSITION_SIDE_BOTH PositionSide = "BOTH"

	// 双向模式下
	POSITION_SIDE_LONG  PositionSide = "LONG"
	POSITION_SIDE_SHORT PositionSide = "SHORT"
)

type OrderType string

const (
	TYPE_LIMIT  OrderType = "LIMIT"
	TYPE_MARKET OrderType = "MARKET"

	// 止盈单
	// 买入: 最新合约价格/标记价格低于等于触发价stopPrice
	// 卖出: 最新合约价格/标记价格高于等于触发价stopPrice
	TYPE_TAKE_PROFIT        OrderType = "TAKE_PROFIT"
	TYPE_TAKE_PROFIT_MARKET OrderType = "TAKE_PROFIT_MARKET"

	// 止损单
	// 买入: 最新合约价格/标记价格高于等于触发价stopPrice
	// 卖出: 最新合约价格/标记价格低于等于触发价stopPrice
	TYPE_STOP        OrderType = "STOP"
	TYPE_STOP_MARKET OrderType = "STOP_MARKET"

	// 跟踪止损单
	// 买入: 当合约价格/标记价格区间最低价格低于激活价格activationPrice,且最新合约价格/标记价高于等于最低价设定回调幅度。
	// 卖出: 当合约价格/标记价格区间最高价格高于激活价格activationPrice,且最新合约价格/标记价低于等于最高价设定回调幅度。
	// TRAILING_STOP_MARKET 跟踪止损单如果遇到报错 {"code": -2021, "msg": "Order would immediately trigger."}
	// 表示订单不满足以下条件:
	// 买入: 指定的activationPrice 必须小于 latest price
	// 卖出: 指定的activationPrice 必须大于 latest price
	TYPE_TRAILING_STOP_MARKET OrderType = "TRAILING_STOP_MARKET"
)

type ReduceOnly string

const (
	REDUCE_ONLY_TRUE  ReduceOnly = "true"
	REDUCE_ONLY_FALSE ReduceOnly = "false"
)

type ClosePosition string

const (
	CLOSE_POSITION_TRUE  ClosePosition = "true"
	CLOSE_POSITION_FALSE ClosePosition = "false"
)

type TimeInForce string

const (
	TIME_IN_FORCE_GTC TimeInForce = "GTC"
	TIME_IN_FORCE_GTD TimeInForce = "GTD"
	TIME_IN_FORCE_IOC TimeInForce = "IOC"
	TIME_IN_FORCE_FOK TimeInForce = "FOK"
)

type WorkingType string

const (
	// 标记价格
	WORKING_TYPE_MARK_PRICE WorkingType = "MARK_PRICE"
	// 最新价格
	WORKING_TYPE_CONTRACT_PRICE WorkingType = "CONTRACT_PRICE"
)

type PriceProtect string

const (
	PRICE_PROTECT_TRUE  PriceProtect = "TRUE"
	PRICE_PROTECT_FALSE PriceProtect = "FALSE"
)

type NewOrderRespType string

const (
	NEW_ORDER_RESP_TYPE_ACK    NewOrderRespType = "ACK"
	NEW_ORDER_RESP_TYPE_RESULT NewOrderRespType = "RESULT"
)

type PriceMatch string

const (
	PRICE_MATCH_NONE       PriceMatch = "NONE"
	PRICE_MATCH_OPPONENT   PriceMatch = "OPPONENT"
	PRICE_MATCH_OPPONENT5  PriceMatch = "OPPONENT5"
	PRICE_MATCH_OPPONENT10 PriceMatch = "OPPONENT10"
	PRICE_MATCH_OPPONENT20 PriceMatch = "OPPONENT20"
	PRICE_MATCH_QUEUE      PriceMatch = "QUEUE"
	PRICE_MATCH_QUEUE5     PriceMatch = "QUEUE5"
	PRICE_MATCH_QUEUE10    PriceMatch = "QUEUE10"
	PRICE_MATCH_QUEUE20    PriceMatch = "QUEUE20"
)

type SelfTradePreventionMode string

const (
	SelfTradePreventionMode_NONE         SelfTradePreventionMode = "NONE"
	SelfTradePreventionMode_EXPIRE_TAKER SelfTradePreventionMode = "EXPIRE_TAKER"
	SelfTradePreventionMode_EXPIRE_MAKER SelfTradePreventionMode = "EXPIRE_MAKER"
	SelfTradePreventionMode_EXPIRE_BOTH  SelfTradePreventionMode = "EXPIRE_BOTH"
)

type OrderStatus string

const (
	ORDER_STATUS_NEW              OrderStatus = "NEW"
	ORDER_STATUS_PARTIALLY_FILLED OrderStatus = "PARTIALLY_FILLED"
	ORDER_STATUS_FILLED           OrderStatus = "FILLED"
	ORDER_STATUS_CANCELED         OrderStatus = "CANCELED"
	ORDER_STATUS_PENDING_CANCEL   OrderStatus = "PENDING_CANCEL"
	ORDER_STATUS_REJECTED         OrderStatus = "REJECTED"
	ORDER_STATUS_EXPIRED          OrderStatus = "EXPIRED"
)
