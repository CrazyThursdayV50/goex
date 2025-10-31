package exchangeinfo

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models"
	"github.com/CrazyThursdayV50/pkgo/json"
)

// ParamsData 交易所信息查询请求参数
type ParamsData struct {
	Symbol  string   `json:"symbol,omitempty"`
	Symbols []string `json:"symbols,omitempty"`

	// https://github.com/binance/binance-spot-api-docs/blob/master/enums_CN.md#account-and-symbol-permissions
	Permissions []string `json:"permissions,omitempty"`

	ShowPermissionSets bool `json:"showPermissionSets,omitempty"`

	// TRADING, HALT, BREAK
	SymbolStatus string `json:"symbolStatus,omitempty"`
}

func NewRequest() *models.WsAPIRequest {
	return &models.WsAPIRequest{
		Method: "exchangeInfo",
	}
}

// SymbolData 交易对信息
type SymbolData struct {
	Symbol                          string     `json:"symbol"`
	Status                          string     `json:"status"`
	BaseAsset                       string     `json:"baseAsset"`
	BaseAssetPrecision              int        `json:"baseAssetPrecision"`
	QuoteAsset                      string     `json:"quoteAsset"`
	QuotePrecision                  int        `json:"quotePrecision"`
	QuoteAssetPrecision             int        `json:"quoteAssetPrecision"`
	BaseCommissionPrecision         int        `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision        int        `json:"quoteCommissionPrecision"`
	OrderTypes                      []string   `json:"orderTypes"`
	IcebergAllowed                  bool       `json:"icebergAllowed"`
	OcoAllowed                      bool       `json:"ocoAllowed"`
	OtoAllowed                      bool       `json:"otoAllowed"`
	QuoteOrderQtyMarketAllowed      bool       `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool       `json:"allowTrailingStop"`
	CancelReplaceAllowed            bool       `json:"cancelReplaceAllowed"`
	AllowAmend                      bool       `json:"allowAmend"`
	IsSpotTradingAllowed            bool       `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool       `json:"isMarginTradingAllowed"`
	Filters                         Filters    `json:"filters"`
	Permissions                     []string   `json:"permissions"`
	PermissionSets                  [][]string `json:"permissionSets"`
	DefaultSelfTradePreventionMode  string     `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string   `json:"allowedSelfTradePreventionModes"`
}

// ResultData 交易所信息查询响应数据
type ResultData struct {
	Timezone        string              `json:"timezone"`
	ServerTime      int64               `json:"serverTime"`
	RateLimits      []*models.RateLimit `json:"rateLimits"`
	ExchangeFilters Filters             `json:"exchangeFilters"`
	Symbols         []*SymbolData       `json:"symbols"`
}

// FilterType 定义过滤器类型
type FilterType string

const (
	PRICE_FILTER                    FilterType = "PRICE_FILTER"
	LOT_SIZE                        FilterType = "LOT_SIZE"
	MIN_NOTIONAL                    FilterType = "MIN_NOTIONAL"
	MARKET_LOT_SIZE                 FilterType = "MARKET_LOT_SIZE"
	MAX_NUM_ORDERS                  FilterType = "MAX_NUM_ORDERS"
	MAX_NUM_ALGO_ORDERS             FilterType = "MAX_NUM_ALGO_ORDERS"
	MAX_NUM_ICEBERG_ORDERS          FilterType = "MAX_NUM_ICEBERG_ORDERS"
	PERCENT_PRICE                   FilterType = "PERCENT_PRICE"
	PERCENT_PRICE_BY_SIDE           FilterType = "PERCENT_PRICE_BY_SIDE"
	ICEBERG_PARTS                   FilterType = "ICEBERG_PARTS"
	MAX_POSITION                    FilterType = "MAX_POSITION"
	TRAILING_DELTA                  FilterType = "TRAILING_DELTA"
	NOTIONAL                        FilterType = "NOTIONAL"
	MAX_OPEN_ORDERS                 FilterType = "MAX_OPEN_ORDERS"
	MAX_OPEN_ALGO_ORDERS            FilterType = "MAX_OPEN_ALGO_ORDERS"
	EXCHANGE_MAX_NUM_ORDERS         FilterType = "EXCHANGE_MAX_NUM_ORDERS"
	EXCHANGE_MAX_ALGO_ORDERS        FilterType = "EXCHANGE_MAX_ALGO_ORDERS"
	EXCHANGE_MAX_NUM_ICEBERG_ORDERS FilterType = "EXCHANGE_MAX_NUM_ICEBERG_ORDERS"
)

// Filter 接口定义所有过滤器必须实现的方法
type Filter interface {
	GetFilterType() FilterType
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

func (f *PriceFilter) MarshalJSON() ([]byte, error) {
	type Alias PriceFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: PRICE_FILTER,
	})
}

func (f *PriceFilter) UnmarshalJSON(data []byte) error {
	type Alias PriceFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != PRICE_FILTER {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
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

func (f *PercentPriceBySideFilter) MarshalJSON() ([]byte, error) {
	type Alias PercentPriceBySideFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: PERCENT_PRICE_BY_SIDE,
	})
}

func (f *PercentPriceBySideFilter) UnmarshalJSON(data []byte) error {
	type Alias PercentPriceBySideFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != PERCENT_PRICE_BY_SIDE {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// LotSizeFilter 数量过滤器
type LotSizeFilter struct {
	BaseFilter
	MinQty   string `json:"minQty"`
	MaxQty   string `json:"maxQty"`
	StepSize string `json:"stepSize"`
}

func (f *LotSizeFilter) MarshalJSON() ([]byte, error) {
	type Alias LotSizeFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: LOT_SIZE,
	})
}

func (f *LotSizeFilter) UnmarshalJSON(data []byte) error {
	type Alias LotSizeFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != LOT_SIZE {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// MinNotionalFilter 最小名义价值过滤器
type MinNotionalFilter struct {
	BaseFilter
	MinNotional      string `json:"minNotional"`
	ApplyToMarket    bool   `json:"applyToMarket"`
	AveragePriceMins int    `json:"avgPriceMins"`
}

func (f *MinNotionalFilter) MarshalJSON() ([]byte, error) {
	type Alias MinNotionalFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: MIN_NOTIONAL,
	})
}

func (f *MinNotionalFilter) UnmarshalJSON(data []byte) error {
	type Alias MinNotionalFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != MIN_NOTIONAL {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// MarketLotSizeFilter 市价订单数量过滤器
type MarketLotSizeFilter struct {
	BaseFilter
	MinQty   string `json:"minQty"`
	MaxQty   string `json:"maxQty"`
	StepSize string `json:"stepSize"`
}

func (f *MarketLotSizeFilter) MarshalJSON() ([]byte, error) {
	type Alias MarketLotSizeFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: MARKET_LOT_SIZE,
	})
}

func (f *MarketLotSizeFilter) UnmarshalJSON(data []byte) error {
	type Alias MarketLotSizeFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != MARKET_LOT_SIZE {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// MaxNumOrdersFilter 最大订单数过滤器
type MaxNumOrdersFilter struct {
	BaseFilter
	MaxNumOrders int `json:"maxNumOrders"`
}

func (f *MaxNumOrdersFilter) MarshalJSON() ([]byte, error) {
	type Alias MaxNumOrdersFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: MAX_NUM_ORDERS,
	})
}

func (f *MaxNumOrdersFilter) UnmarshalJSON(data []byte) error {
	type Alias MaxNumOrdersFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != MAX_NUM_ORDERS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// MaxNumAlgoOrdersFilter 最大算法订单数过滤器
type MaxNumAlgoOrdersFilter struct {
	BaseFilter
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders"`
}

func (f *MaxNumAlgoOrdersFilter) MarshalJSON() ([]byte, error) {
	type Alias MaxNumAlgoOrdersFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: MAX_NUM_ALGO_ORDERS,
	})
}

func (f *MaxNumAlgoOrdersFilter) UnmarshalJSON(data []byte) error {
	type Alias MaxNumAlgoOrdersFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != MAX_NUM_ALGO_ORDERS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// PercentPriceFilter 价格振幅过滤器
type PercentPriceFilter struct {
	BaseFilter
	MultiplierUp   string `json:"multiplierUp"`
	MultiplierDown string `json:"multiplierDown"`
	AvgPriceMins   int    `json:"avgPriceMins"`
}

func (f *PercentPriceFilter) MarshalJSON() ([]byte, error) {
	type Alias PercentPriceFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: PERCENT_PRICE,
	})
}

func (f *PercentPriceFilter) UnmarshalJSON(data []byte) error {
	type Alias PercentPriceFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != PERCENT_PRICE {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// IcebergPartsFilter 冰山订单拆分数过滤器
type IcebergPartsFilter struct {
	BaseFilter
	Limit int `json:"limit"`
}

func (f *IcebergPartsFilter) MarshalJSON() ([]byte, error) {
	type Alias IcebergPartsFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: ICEBERG_PARTS,
	})
}

func (f *IcebergPartsFilter) UnmarshalJSON(data []byte) error {
	type Alias IcebergPartsFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != ICEBERG_PARTS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// MaxNumIcebergOrdersFilter 最大冰山订单数过滤器
type MaxNumIcebergOrdersFilter struct {
	BaseFilter
	MaxNumIcebergOrders int `json:"maxNumIcebergOrders"`
}

func (f *MaxNumIcebergOrdersFilter) MarshalJSON() ([]byte, error) {
	type Alias MaxNumIcebergOrdersFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: MAX_NUM_ICEBERG_ORDERS,
	})
}

func (f *MaxNumIcebergOrdersFilter) UnmarshalJSON(data []byte) error {
	type Alias MaxNumIcebergOrdersFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != MAX_NUM_ICEBERG_ORDERS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// MaxPositionRiskFilter 最大持仓风险过滤器
type MaxPositionRiskFilter struct {
	BaseFilter
	Limit int `json:"limit"`
}

func (f *MaxPositionRiskFilter) MarshalJSON() ([]byte, error) {
	type Alias MaxPositionRiskFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: MAX_POSITION,
	})
}

func (f *MaxPositionRiskFilter) UnmarshalJSON(data []byte) error {
	type Alias MaxPositionRiskFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != MAX_POSITION {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// TrailingDeltaFilter 追踪止损过滤器
type TrailingDeltaFilter struct {
	BaseFilter
	MinTrailingAboveDelta int `json:"minTrailingAboveDelta"`
	MaxTrailingAboveDelta int `json:"maxTrailingAboveDelta"`
	MinTrailingBelowDelta int `json:"minTrailingBelowDelta"`
	MaxTrailingBelowDelta int `json:"maxTrailingBelowDelta"`
}

func (f *TrailingDeltaFilter) MarshalJSON() ([]byte, error) {
	type Alias TrailingDeltaFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: TRAILING_DELTA,
	})
}

func (f *TrailingDeltaFilter) UnmarshalJSON(data []byte) error {
	type Alias TrailingDeltaFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != TRAILING_DELTA {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
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

func (f *NotionalFilter) MarshalJSON() ([]byte, error) {
	type Alias NotionalFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: NOTIONAL,
	})
}

func (f *NotionalFilter) UnmarshalJSON(data []byte) error {
	type Alias NotionalFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != NOTIONAL {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// MaxOpenOrdersFilter 最大未成交订单数过滤器
type MaxOpenOrdersFilter struct {
	BaseFilter
	MaxOpenOrders int `json:"maxOpenOrders"`
}

func (f *MaxOpenOrdersFilter) MarshalJSON() ([]byte, error) {
	type Alias MaxOpenOrdersFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: MAX_OPEN_ORDERS,
	})
}

func (f *MaxOpenOrdersFilter) UnmarshalJSON(data []byte) error {
	type Alias MaxOpenOrdersFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != MAX_OPEN_ORDERS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// MaxOpenAlgoOrdersFilter 最大未成交算法订单数过滤器
type MaxOpenAlgoOrdersFilter struct {
	BaseFilter
	MaxOpenAlgoOrders int `json:"maxOpenAlgoOrders"`
}

func (f *MaxOpenAlgoOrdersFilter) MarshalJSON() ([]byte, error) {
	type Alias MaxOpenAlgoOrdersFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: MAX_OPEN_ALGO_ORDERS,
	})
}

func (f *MaxOpenAlgoOrdersFilter) UnmarshalJSON(data []byte) error {
	type Alias MaxOpenAlgoOrdersFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != MAX_OPEN_ALGO_ORDERS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// ExchangeMaxNumOrdersFilter 交易所最大订单数过滤器
type ExchangeMaxNumOrdersFilter struct {
	BaseFilter
	MaxNumOrders int `json:"maxNumOrders"`
}

func (f *ExchangeMaxNumOrdersFilter) MarshalJSON() ([]byte, error) {
	type Alias ExchangeMaxNumOrdersFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: EXCHANGE_MAX_NUM_ORDERS,
	})
}

func (f *ExchangeMaxNumOrdersFilter) UnmarshalJSON(data []byte) error {
	type Alias ExchangeMaxNumOrdersFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != EXCHANGE_MAX_NUM_ORDERS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// ExchangeMaxAlgoOrdersFilter 交易所最大算法订单数过滤器
type ExchangeMaxAlgoOrdersFilter struct {
	BaseFilter
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders"`
}

func (f *ExchangeMaxAlgoOrdersFilter) MarshalJSON() ([]byte, error) {
	type Alias ExchangeMaxAlgoOrdersFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: EXCHANGE_MAX_ALGO_ORDERS,
	})
}

func (f *ExchangeMaxAlgoOrdersFilter) UnmarshalJSON(data []byte) error {
	type Alias ExchangeMaxAlgoOrdersFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != EXCHANGE_MAX_ALGO_ORDERS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// ExchangeMaxNumIcebergOrdersFilter 交易所最大冰山订单数过滤器
type ExchangeMaxNumIcebergOrdersFilter struct {
	BaseFilter
	MaxNumIcebergOrders int `json:"maxNumIcebergOrders"`
}

func (f *ExchangeMaxNumIcebergOrdersFilter) MarshalJSON() ([]byte, error) {
	type Alias ExchangeMaxNumIcebergOrdersFilter
	return json.JSON().Marshal(&struct {
		*Alias
		FilterType FilterType `json:"filterType"`
	}{
		Alias:      (*Alias)(f),
		FilterType: EXCHANGE_MAX_NUM_ICEBERG_ORDERS,
	})
}

func (f *ExchangeMaxNumIcebergOrdersFilter) UnmarshalJSON(data []byte) error {
	type Alias ExchangeMaxNumIcebergOrdersFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.JSON().Unmarshal(data, aux); err != nil {
		return err
	}
	if f.FilterType != EXCHANGE_MAX_NUM_ICEBERG_ORDERS {
		return fmt.Errorf("invalid filter type: %s", f.FilterType)
	}
	return nil
}

// GenericFilter 通用过滤器，用于处理未知的过滤器类型
type GenericFilter struct {
	BaseFilter
	RawData map[string]any `json:"-"`
}

func (f *GenericFilter) MarshalJSON() ([]byte, error) {
	if f.RawData == nil {
		f.RawData = make(map[string]any)
	}
	f.RawData["filterType"] = f.FilterType
	return json.JSON().Marshal(f.RawData)
}

func (f *GenericFilter) UnmarshalJSON(data []byte) error {
	// 首先解析基础过滤器类型
	if err := json.JSON().Unmarshal(data, &f.BaseFilter); err != nil {
		return err
	}

	// 解析原始数据到 map
	if err := json.JSON().Unmarshal(data, &f.RawData); err != nil {
		return err
	}

	return nil
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
