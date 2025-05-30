package order

// OrderSide 订单方向
type OrderSide string

const (
	BUY  OrderSide = "BUY"
	SELL OrderSide = "SELL"
)

type OrderStatus string

const (
	NEW              OrderStatus = "NEW"
	PENDING_NEW      OrderStatus = "PENDING_NEW"
	PARTIALLY_FILLED OrderStatus = "PARTIALLY_FILLED"
	FILLED           OrderStatus = "FILLED"
	CANCELED         OrderStatus = "CANCELED"
	PENDING_CANCEL   OrderStatus = "PENDING_CANCEL"
	REJECTED         OrderStatus = "REJECTED"
	EXPIRED          OrderStatus = "EXPIRED"
	EXPIRED_IN_MATCH OrderStatus = "EXPIRED_IN_MATCH"
)

// OrderType 订单类型
type OrderType string

const (
	LIMIT             OrderType = "LIMIT"
	MARKET            OrderType = "MARKET"
	STOP_LOSS         OrderType = "STOP_LOSS"
	STOP_LOSS_LIMIT   OrderType = "STOP_LOSS_LIMIT"
	TAKE_PROFIT       OrderType = "TAKE_PROFIT"
	TAKE_PROFIT_LIMIT OrderType = "TAKE_PROFIT_LIMIT"
	LIMIT_MAKER       OrderType = "LIMIT_MAKER"
)

// TimeInForce 订单有效期
type TimeInForce string

const (
	GTC TimeInForce = "GTC" // Good Till Cancel
	IOC TimeInForce = "IOC" // Immediate or Cancel
	FOK TimeInForce = "FOK" // Fill or Kill
)

type PegPriceType string

const (
	PRIMARY_PEG PegPriceType = "PRIMARY_PEG"
	MARKET_PEG  PegPriceType = "MARKET_PED"
)

type PegOffsetType string

const (
	PRICE_LEVEL PegOffsetType = "PRICE_LEVEL"
)

type OrderRespType string

const (
	ACK    OrderRespType = "ACK"
	RESULT OrderRespType = "RESULT"
	FULL   OrderRespType = "FULL"
)

type SelfTradePreventionMode string

const (
	NONE         SelfTradePreventionMode = "NONE"
	EXPIRE_MAKER SelfTradePreventionMode = "EXPIRE_MAKER"
	EXPIRE_TAKER SelfTradePreventionMode = "EXPIRE_TAKER"
	EXPIRE_BOTH  SelfTradePreventionMode = "EXPIRE_BOTH"
	DECREMENT    SelfTradePreventionMode = "DECREMENT"
)

type ListStatusType string

const (
	// 在 ListStatus 用于响应失败的操作时会被使用。（例如，下订单组或取消订单组）
	ListStatus_RESPONSE ListStatusType = "RESPONSE"
	// 订单组已被下达或订单组状态有更新。
	ListStatus_EXEC_STARTED ListStatusType = "EXEC_STARTED"
	// 订单组里的某个订单的 clientOrderId 被改变。
	ListStatus_UPDATED ListStatusType = "UPDATED"
	// 订单组执行结束，因此不再处于活动状态。
	ListStatus_ALL_DONE ListStatusType = "ALL_DONE"
)

type ListOrderStatus string

const (
	// 订单组已被下达或订单组状态有更新。
	ListOrder_EXECUTING ListOrderStatus = "EXEC_STARTED"
	// 订单组执行结束，因此不再处于活动状态。
	ListOrder_ALL_DONE ListOrderStatus = "ALL_DONE"
	// 在 ListStatus 用于响应在下单阶段或取消订单组期间的失败操作时会被使用，
	ListOrder_REJECT ListOrderStatus = "UPDATED"
)

type ListType string

const (
	OCO ListType = "OCO"
	OTO ListType = "OTO"
)
