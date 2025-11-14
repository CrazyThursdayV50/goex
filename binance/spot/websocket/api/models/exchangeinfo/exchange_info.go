package exchangeinfo

import (
	"github.com/CrazyThursdayV50/goex/binance/models/filters"
	"github.com/CrazyThursdayV50/goex/binance/models/ratelimits"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models"
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
	Symbol                          string          `json:"symbol"`
	Status                          string          `json:"status"`
	BaseAsset                       string          `json:"baseAsset"`
	BaseAssetPrecision              int             `json:"baseAssetPrecision"`
	QuoteAsset                      string          `json:"quoteAsset"`
	QuotePrecision                  int             `json:"quotePrecision"`
	QuoteAssetPrecision             int             `json:"quoteAssetPrecision"`
	BaseCommissionPrecision         int             `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision        int             `json:"quoteCommissionPrecision"`
	OrderTypes                      []string        `json:"orderTypes"`
	IcebergAllowed                  bool            `json:"icebergAllowed"`
	OcoAllowed                      bool            `json:"ocoAllowed"`
	OtoAllowed                      bool            `json:"otoAllowed"`
	QuoteOrderQtyMarketAllowed      bool            `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool            `json:"allowTrailingStop"`
	CancelReplaceAllowed            bool            `json:"cancelReplaceAllowed"`
	AllowAmend                      bool            `json:"allowAmend"`
	IsSpotTradingAllowed            bool            `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool            `json:"isMarginTradingAllowed"`
	Filters                         filters.Filters `json:"filters"`
	Permissions                     []string        `json:"permissions"`
	PermissionSets                  [][]string      `json:"permissionSets"`
	DefaultSelfTradePreventionMode  string          `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string        `json:"allowedSelfTradePreventionModes"`
}

// ResultData 交易所信息查询响应数据
type ResultData struct {
	Timezone        string                 `json:"timezone"`
	ServerTime      int64                  `json:"serverTime"`
	RateLimits      []ratelimits.RateLimit `json:"rateLimits"`
	ExchangeFilters filters.Filters        `json:"exchangeFilters"`
	Symbols         []*SymbolData          `json:"symbols"`
}
