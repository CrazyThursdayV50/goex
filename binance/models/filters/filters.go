package filters

import (
	"fmt"

	"github.com/CrazyThursdayV50/pkgo/json"
)

// Filter 接口定义所有过滤器必须实现的方法
type Filter interface {
	GetFilterType() FilterType
}

// Filters 定义一个可以容纳不同类型 Filter 的切片
type Filters []Filter

// UnmarshalJSON 实现 Filters 类型的 UnmarshalJSON 方法
func (f *Filters) UnmarshalJSON(data []byte) error {
	// data 是包含 Filter JSON 对象的数组
	var rawFilters []json.RawMessage
	if err := json.JSON().Unmarshal(data, &rawFilters); err != nil {
		return err
	}

	// 遍历原始的 JSON 消息，并解析成具体的 Filter 类型
	parsedFilters := make(Filters, 0, len(rawFilters))
	for _, rawFilter := range rawFilters {
		filter, err := ParseFilter(rawFilter)
		if err != nil {
			return err
		}
		parsedFilters = append(parsedFilters, filter)
	}

	*f = parsedFilters
	return nil
}

// MarshalJSON 实现 Filters 类型的 MarshalJSON 方法
func (f Filters) MarshalJSON() ([]byte, error) {
	return json.JSON().Marshal([]Filter(f))
}

// BaseFilter 所有过滤器的基础结构
type BaseFilter struct {
	FilterType FilterType `json:"filterType"`
}

func (f BaseFilter) GetFilterType() FilterType {
	return f.FilterType
}

// PriceFilter 价格过滤器
type PriceFilter struct {
	BaseFilter
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice"`
	TickSize string `json:"tickSize"`
}

// PercentPriceBySideFilter 按买卖方向的价格振幅过滤器
type PercentPriceBySideFilter struct {
	BaseFilter
	BidMultiplierUp   string `json:"bidMultiplierUp"`
	BidMultiplierDown string `json:"bidMultiplierDown"`
	AskMultiplierUp   string `json:"askMultiplierUp"`
	AskMultiplierDown string `json:"askMultiplierDown"`
	AvgPriceMins      int    `json:"avgPriceMins"`
}

// LotSizeFilter 数量过滤器
type LotSizeFilter struct {
	BaseFilter
	MinQty   string `json:"minQty"`
	MaxQty   string `json:"maxQty"`
	StepSize string `json:"stepSize"`
}

// MinNotionalFilter 最小名义价值过滤器
type MinNotionalFilter struct {
	BaseFilter
	MinNotional      string `json:"minNotional"`
	ApplyToMarket    bool   `json:"applyToMarket"`
	AveragePriceMins int    `json:"avgPriceMins"`
}

// MarketLotSizeFilter 市价订单数量过滤器
type MarketLotSizeFilter struct {
	BaseFilter
	MinQty   string `json:"minQty"`
	MaxQty   string `json:"maxQty"`
	StepSize string `json:"stepSize"`
}

// MaxNumOrdersFilter 最大订单数过滤器
type MaxNumOrdersFilter struct {
	BaseFilter
	MaxNumOrders int `json:"maxNumOrders"`
}

// MaxNumAlgoOrdersFilter 最大算法订单数过滤器
type MaxNumAlgoOrdersFilter struct {
	BaseFilter
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders"`
}

// PercentPriceFilter 价格振幅过滤器
type PercentPriceFilter struct {
	BaseFilter
	MultiplierUp   string `json:"multiplierUp"`
	MultiplierDown string `json:"multiplierDown"`
	AvgPriceMins   int    `json:"avgPriceMins"`
}

// IcebergPartsFilter 冰山订单拆分数过滤器
type IcebergPartsFilter struct {
	BaseFilter
	Limit int `json:"limit"`
}

// MaxNumIcebergOrdersFilter 最大冰山订单数过滤器
type MaxNumIcebergOrdersFilter struct {
	BaseFilter
	MaxNumIcebergOrders int `json:"maxNumIcebergOrders"`
}

// MaxPositionRiskFilter 最大持仓风险过滤器
type MaxPositionRiskFilter struct {
	BaseFilter
	Limit int `json:"limit"`
}

// TrailingDeltaFilter 追踪止损过滤器
type TrailingDeltaFilter struct {
	BaseFilter
	MinTrailingAboveDelta int `json:"minTrailingAboveDelta"`
	MaxTrailingAboveDelta int `json:"maxTrailingAboveDelta"`
	MinTrailingBelowDelta int `json:"minTrailingBelowDelta"`
	MaxTrailingBelowDelta int `json:"maxTrailingBelowDelta"`
}

// NotionalFilter 名义价值过滤器
type NotionalFilter struct {
	BaseFilter
	MinNotional           string `json:"minNotional"`
	ApplyMinToMarket      bool   `json:"applyMinToMarket"`
	MaxNotional           string `json:"maxNotional"`
	ApplyMaxToMarket      bool   `json:"applyMaxToMarket"`
	AveragePriceMins      int    `json:"avgPriceMins"`
	CalculateNotionalType string `json:"calculateNotionalType"`
}

// MaxOpenOrdersFilter 最大未成交订单数过滤器
type MaxOpenOrdersFilter struct {
	BaseFilter
	MaxOpenOrders int `json:"maxOpenOrders"`
}

// MaxOpenAlgoOrdersFilter 最大未成交算法订单数过滤器
type MaxOpenAlgoOrdersFilter struct {
	BaseFilter
	MaxOpenAlgoOrders int `json:"maxOpenAlgoOrders"`
}

// ExchangeMaxNumOrdersFilter 交易所最大订单数过滤器
type ExchangeMaxNumOrdersFilter struct {
	BaseFilter
	MaxNumOrders int `json:"maxNumOrders"`
}

// ExchangeMaxAlgoOrdersFilter 交易所最大算法订单数过滤器
type ExchangeMaxAlgoOrdersFilter struct {
	BaseFilter
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders"`
}

// ExchangeMaxNumIcebergOrdersFilter 交易所最大冰山订单数过滤器
type ExchangeMaxNumIcebergOrdersFilter struct {
	BaseFilter
	MaxNumIcebergOrders int `json:"maxNumIcebergOrders"`
}

// GenericFilter 通用过滤器，用于处理未知的过滤器类型
type GenericFilter struct {
	BaseFilter
	RawData map[string]any `json:"-"`
}

// ParseFilter 解析过滤器 JSON 数据
func ParseFilter(data []byte) (Filter, error) {
	var base BaseFilter
	if err := json.JSON().Unmarshal(data, &base); err != nil {
		return nil, err
	}

	var filter Filter
	switch base.FilterType {
	case PRICE_FILTER:
		filter = &PriceFilter{}
	case LOT_SIZE:
		filter = &LotSizeFilter{}
	case MIN_NOTIONAL:
		filter = &MinNotionalFilter{}
	case MARKET_LOT_SIZE:
		filter = &MarketLotSizeFilter{}
	case MAX_NUM_ORDERS:
		filter = &MaxNumOrdersFilter{}
	case MAX_NUM_ALGO_ORDERS:
		filter = &MaxNumAlgoOrdersFilter{}
	case PERCENT_PRICE:
		filter = &PercentPriceFilter{}
	case PERCENT_PRICE_BY_SIDE:
		filter = &PercentPriceBySideFilter{}
	case ICEBERG_PARTS:
		filter = &IcebergPartsFilter{}
	case MAX_NUM_ICEBERG_ORDERS:
		filter = &MaxNumIcebergOrdersFilter{}
	case MAX_POSITION:
		filter = &MaxPositionRiskFilter{}
	case TRAILING_DELTA:
		filter = &TrailingDeltaFilter{}
	case NOTIONAL:
		filter = &NotionalFilter{}
	case MAX_OPEN_ORDERS:
		filter = &MaxOpenOrdersFilter{}
	case MAX_OPEN_ALGO_ORDERS:
		filter = &MaxOpenAlgoOrdersFilter{}
	case EXCHANGE_MAX_NUM_ORDERS:
		filter = &ExchangeMaxNumOrdersFilter{}
	case EXCHANGE_MAX_ALGO_ORDERS:
		filter = &ExchangeMaxAlgoOrdersFilter{}
	case EXCHANGE_MAX_NUM_ICEBERG_ORDERS:
		filter = &ExchangeMaxNumIcebergOrdersFilter{}
	default:
		// 使用通用过滤器处理未知类型
		genericFilter := &GenericFilter{}
		if err := json.JSON().Unmarshal(data, genericFilter); err != nil {
			return nil, fmt.Errorf("failed to parse unknown filter type %s: %v", base.FilterType, err)
		}
		return genericFilter, nil
	}

	if err := json.JSON().Unmarshal(data, filter); err != nil {
		return nil, err
	}

	return filter, nil
}
