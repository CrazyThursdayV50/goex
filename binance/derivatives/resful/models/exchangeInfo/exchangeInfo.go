package exchangeinfo

import (
	"github.com/CrazyThursdayV50/goex/binance/models/filters"
	"github.com/CrazyThursdayV50/goex/binance/models/ratelimits"
)

type Result struct {
	ServerTime int64                  `json:"serverTime"` // 服务器时间 (毫秒时间戳)
	RateLimits []ratelimits.RateLimit `json:"rateLimits"`
	Assets     []AssetInfo            `json:"assets"`
	Symbols    []SymbolInfo           `json:"symbols"`
	Timezone   string                 `json:"timezone"`
}

type AssetInfo struct {
	// 资产名称，如 BTC, USDT
	Asset string `json:"asset"`

	// 是否可用作保证金
	MarginAvailable bool `json:"marginAvailable"`

	// 保证金资产自动兑换阈值。注意：此字段可能返回字符串 ("-0.10", "0") 或 null，
	// 因此推荐使用 *string 指针以安全地处理 null 值，或者使用 interface{}。
	// 但对于大多数情况，如果 API 文档确认 null 出现频率高且需要区分，使用 *string 更佳。
	// 如果 null 很少出现且 "0" 含义明确，也可以直接使用 string。
	// 这里使用 *string 来演示对 null 的处理。
	AutoAssetExchange *string `json:"autoAssetExchange"`
}

type SymbolInfo struct {
	Symbol                string          `json:"symbol"`                // 交易对
	Pair                  string          `json:"pair"`                  // 标的交易对
	ContractType          string          `json:"contractType"`          // 合约类型
	DeliveryDate          int64           `json:"deliveryDate"`          // 交割日期 (毫秒时间戳)
	OnboardDate           int64           `json:"onboardDate"`           // 上线日期 (毫秒时间戳)
	Status                string          `json:"status"`                // 交易对状态
	MaintMarginPercent    string          `json:"maintMarginPercent"`    // 维持保证金率 (忽略，但保持以便完整解析)
	RequiredMarginPercent string          `json:"requiredMarginPercent"` // 所需初始保证金率 (忽略，但保持以便完整解析)
	BaseAsset             string          `json:"baseAsset"`             // 标的资产 (如 BLZ)
	QuoteAsset            string          `json:"quoteAsset"`            // 报价资产 (如 USDT)
	MarginAsset           string          `json:"marginAsset"`           // 保证金资产 (如 USDT)
	PricePrecision        int64           `json:"pricePrecision"`        // 价格小数点位数
	QuantityPrecision     int64           `json:"quantityPrecision"`     // 数量小数点位数
	BaseAssetPrecision    int64           `json:"baseAssetPrecision"`    // 标的资产精度
	QuotePrecision        int64           `json:"quotePrecision"`        // 报价资产精度
	UnderlyingType        string          `json:"underlyingType"`
	UnderlyingSubType     []string        `json:"underlyingSubType"`
	SettlePlan            int64           `json:"settlePlan"`
	TriggerProtect        string          `json:"triggerProtect"`
	Filters               filters.Filters `json:"filters"`
	OrderType             []string        `json:"OrderType"`
	TimeInForce           []string        `json:"timeInForce"`
	LiquidationFee        string          `json:"liquidationFee"`  // 强平费率
	MarketTakeBound       string          `json:"marketTakeBound"` // 市价吃单(相对于标记价格)允许可造成的最大价格偏离比例
}
