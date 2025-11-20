package order

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models"

// 当有新订单创建、订单有新成交或者新的状态变化时会推送此类事件 事件类型统一为 ORDER_TRADE_UPDATE
const (
	EventName = "ORDER_TRADE_UPDATE"
)

type Event struct {
	models.BaseEvent
	Data EventData `json:"o"`
}

// OrderUpdateEvent 订单更新事件结构体
type EventData struct {
	// 交易对
	Symbol string `json:"s"`
	// 客户端自定订单ID
	ClientOrderID string `json:"c"`
	// 订单方向
	Side OrderSide `json:"S"`
	// 订单类型
	Type OrderType `json:"o"`
	// 有效方式
	TimeInForce TimeInForceType `json:"f"`
	// 订单原始数量 (使用 string 保持精度)
	OriginalQuantity string `json:"q"`
	// 订单原始价格 (使用 string 保持精度)
	OriginalPrice string `json:"p"`
	// 订单平均价格 (使用 string 保持精度)
	AveragePrice string `json:"ap"`
	// 条件订单触发价格 (对追踪止损单无效，使用 string 保持精度)
	StopPrice string `json:"sp"`
	// 本次事件的具体执行类型
	ExecutionType ExecutionType `json:"x"`
	// 订单的当前状态
	Status OrderStatus `json:"X"`
	// 订单ID
	OrderID int64 `json:"i"`
	// 订单末次成交量 (使用 string 保持精度)
	LastExecutedQuantity string `json:"l"`
	// 订单累计已成交量 (使用 string 保持精度)
	CumulativeFilledQuantity string `json:"z"`
	// 订单末次成交价格 (使用 string 保持精度)
	LastExecutedPrice string `json:"L"`
	// 手续费资产类型
	CommissionAsset string `json:"N"`
	// 手续费数量 (使用 string 保持精度)
	Commission string `json:"n"`
	// 成交时间 (Unix 毫秒时间戳)
	TradeTime int64 `json:"T"`
	// 成交ID
	TradeID int64 `json:"t"`
	// 买单净值 (使用 string 保持精度)
	BuyerCommission string `json:"b"`
	// 卖单净值 (使用 string 保持精度)
	SellerCommission string `json:"a"`
	// 该成交是作为挂单成交吗？ (true: 挂单(Maker), false: 吃单(Taker))
	IsMaker bool `json:"m"`
	// 是否是只减仓单
	IsReduceOnly bool `json:"R"`
	// 触发价类型 (e.g., "CONTRACT_PRICE")
	TriggerPriceType string `json:"wt"`
	// 原始订单类型
	OriginalOrderType OrderType `json:"ot"`
	// 持仓方向 (LONG, SHORT, BOTH)
	PositionSide string `json:"ps"`
	// 是否为触发平仓单; 仅在条件订单情况下会推送此字段
	IsTriggeredCloseOrder bool `json:"cp"`
	// 追踪止损激活价格 (仅在追踪止损单时推送，使用 string 保持精度)
	ActivationPrice string `json:"AP"`
	// 追踪止损回调比例 (仅在追踪止损单时推送，使用 string 保持精度)
	CallbackRate string `json:"cr"`
	// 是否开启条件单触发保护
	IsProtectionEnabled bool `json:"pP"`
	// 忽略
	Ignore1 int `json:"si"`
	// 忽略
	Ignore2 int `json:"ss"`
	// 该交易实现盈亏 (使用 string 保持精度)
	RealizedProfitLoss string `json:"rp"`
	// 自成交防止模式 (e.g., "EXPIRE_TAKER")
	SelfTradePreventionMode string `json:"V"`
	// 价格匹配模式 (e.g., "OPPONENT")
	PriceMatchingMode string `json:"pm"`
	// TIF为GTD的订单自动取消时间 (Unix 毫秒时间戳)
	GoodTillDate int64 `json:"gtd"`
	// 过期原因
	ExpireReason ExpireReasonCode `json:"er"`
}
