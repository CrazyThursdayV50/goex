package models

import (
	"errors"
	"fmt"

	"github.com/CrazyThursdayV50/pkgo/json"
)

type WsAPIParams[T any] struct {
	Id     string `json:"id"`
	Method string `json:"method"`
	Params T      `json:"params,omitempty"`
}

func (p *WsAPIParams[T]) BinaryMarshal() ([]byte, error) {
	return json.JSON().Marshal(p)
}

func (p *WsAPIParams[T]) BinaryUnmarshal(data []byte) error {
	if p == nil {
		return errors.New("nil receiver")
	}

	return json.JSON().Unmarshal(data, p)
}

type WsAPIResult struct {
	Id         string            `json:"id"`
	Status     int               `json:"status"`
	Result     json.RawMessage   `json:"result"`
	RateLimits []*WsAPIRateLimit `json:"rateLimits"`
}

func (r *WsAPIResult) String() string {
	return fmt.Sprintf("[%s] - %d - %s", r.Id, r.Status, r.Result)
}

type WsExchangeInfoParamsData struct {
	Symbol  string   `json:"symbol,omitempty"`
	Symbols []string `json:"symbols,omitempty"`

	// https://github.com/binance/binance-spot-api-docs/blob/master/enums_CN.md#account-and-symbol-permissions
	Permissions []string `json:"permissions,omitempty"`

	ShowPermissionSets bool `json:"showPermissionSets,omitempty"`

	// TRADING, HALT, BREAK
	SymbolStatus string `json:"symbolStatus,omitempty"`
}

/*
	{
	  "symbol": "BNBBTC",
	  "status": "TRADING",
	  "baseAsset": "BNB",
	  "baseAssetPrecision": 8,
	  "quoteAsset": "BTC",
	  "quotePrecision": 8,
	  "quoteAssetPrecision": 8,
	  "baseCommissionPrecision": 8,
	  "quoteCommissionPrecision": 8,
	  "orderTypes": [
	    "LIMIT",
	    "LIMIT_MAKER",
	    "MARKET",
	    "STOP_LOSS_LIMIT",
	    "TAKE_PROFIT_LIMIT"
	  ],
	  "icebergAllowed": true,
	  "ocoAllowed": true,
	  "otoAllowed": true,
	  "quoteOrderQtyMarketAllowed": true,
	  "allowTrailingStop": true,
	  "cancelReplaceAllowed": true,
	  "allowAmend":false,
	  "isSpotTradingAllowed": true,
	  "isMarginTradingAllowed": true,
	  // 交易对过滤器在"过滤器"页面上进行了说明：
	  // https://github.com/binance/binance-spot-api-docs/blob/master/filters_CN.md
	  // 全部交易对过滤器是可选的。
	  "filters": [
	    {
	      "filterType": "PRICE_FILTER",
	      "minPrice": "0.00000100",
	      "maxPrice": "100000.00000000",
	      "tickSize": "0.00000100"
	    },
	    {
	      "filterType": "LOT_SIZE",
	      "minQty": "0.00100000",
	      "maxQty": "100000.00000000",
	      "stepSize": "0.00100000"
	    }
	  ],
	  "permissions": [],
	  "permissionSets": [
	    [
	      "SPOT",
	      "MARGIN",
	      "TRD_GRP_004"
	    ]
	  ],
	  "defaultSelfTradePreventionMode": "NONE",
	  "allowedSelfTradePreventionModes": [
	    "NONE"
	  ]
	}
*/

type WsExchangeInfoSymbolData struct {
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

type WsExchangeInfoResultData struct {
	Timezone        string                      `json:"timezone"`
	ServerTime      int64                       `json:"serverTime"`
	RateLimits      []*WsAPIRateLimit           `json:"rateLimits"`
	ExchangeFilters Filters                     `json:"exchangeFilters"`
	Symbols         []*WsExchangeInfoSymbolData `json:"symbols"`
	// Sors            []Sor                       `json:"sors,omitempty"`
}

/*
{
  "id": "5494febb-d167-46a2-996d-70533eb4d976",
  "status": 200,
  "result": {
    "timezone": "UTC",
    "serverTime": 1655969291181,
    // 全局速率限制。请参阅 "速率限制" 部分。
    "rateLimits": [
      {
        "rateLimitType": "REQUEST_WEIGHT",    // 速率限制类型: REQUEST_WEIGHT，ORDERS，CONNECTIONS
        "interval": "MINUTE",                 // 速率限制间隔: SECOND，MINUTE，DAY
        "intervalNum": 1,                     // 速率限制间隔乘数 (i.e.，"1 minute")
        "limit": 6000                         // 每个间隔的速率限制
      },
      {
        "rateLimitType": "ORDERS",
        "interval": "SECOND",
        "intervalNum": 10,
        "limit": 50
      },
      {
        "rateLimitType": "ORDERS",
        "interval": "DAY",
        "intervalNum": 1,
        "limit": 160000
      },
      {
        "rateLimitType": "CONNECTIONS",
        "interval": "MINUTE",
        "intervalNum": 5,
        "limit": 300
      }
    ],
    // 交易所级别过滤器在 "过滤器" 页面上进行了说明：
    // https://github.com/binance/binance-spot-api-docs/blob/master/filters_CN.md
    // 全部交易过滤器是可选的。
    "exchangeFilters": [],
    "symbols": [
      {
        "symbol": "BNBBTC",
        "status": "TRADING",
        "baseAsset": "BNB",
        "baseAssetPrecision": 8,
        "quoteAsset": "BTC",
        "quotePrecision": 8,
        "quoteAssetPrecision": 8,
        "baseCommissionPrecision": 8,
        "quoteCommissionPrecision": 8,
        "orderTypes": [
          "LIMIT",
          "LIMIT_MAKER",
          "MARKET",
          "STOP_LOSS_LIMIT",
          "TAKE_PROFIT_LIMIT"
        ],
        "icebergAllowed": true,
        "ocoAllowed": true,
        "otoAllowed": true,
        "quoteOrderQtyMarketAllowed": true,
        "allowTrailingStop": true,
        "cancelReplaceAllowed": true,
        "allowAmend":false,
        "isSpotTradingAllowed": true,
        "isMarginTradingAllowed": true,
        // 交易对过滤器在"过滤器"页面上进行了说明：
        // https://github.com/binance/binance-spot-api-docs/blob/master/filters_CN.md
        // 全部交易对过滤器是可选的。
        "filters": [
          {
            "filterType": "PRICE_FILTER",
            "minPrice": "0.00000100",
            "maxPrice": "100000.00000000",
            "tickSize": "0.00000100"
          },
          {
            "filterType": "LOT_SIZE",
            "minQty": "0.00100000",
            "maxQty": "100000.00000000",
            "stepSize": "0.00100000"
          }
        ],
        "permissions": [],
        "permissionSets": [
          [
            "SPOT",
            "MARGIN",
            "TRD_GRP_004"
          ]
        ],
        "defaultSelfTradePreventionMode": "NONE",
        "allowedSelfTradePreventionModes": [
          "NONE"
        ]
      }
    ],
    // 可选字段，仅当 SOR 可用时才会被显示出来。
    // https://github.com/binance/binance-spot-api-docs/blob/master/faqs/sor_faq_CN.md
    "sors": [
      {
        "baseAsset": "BTC",
        "symbols": [
          "BTCUSDT",
          "BTCUSDC"
        ]
      }
    ]
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 20
    }
  ]
}
*/

type WsPingParams = WsAPIParams[any]

func NewWsPingParams() *WsPingParams {
	return &WsPingParams{Method: "ping"}
}

type WsExchangeInfoParams = WsAPIParams[*WsExchangeInfoParamsData]

func NewWsExchangeInfoParams() *WsExchangeInfoParams {
	return &WsExchangeInfoParams{Method: "exchangeInfo", Params: &WsExchangeInfoParamsData{}}
}

/*
	{
	  "id": "922bcc6e-9de8-440d-9e84-7c80933a8d0d",
	  "status": 200,
	  "result": {},
	  "rateLimits": [
	    {
	      "rateLimitType": "REQUEST_WEIGHT",
	      "interval": "MINUTE",
	      "intervalNum": 1,
	      "limit": 6000,
	      "count": 1
	    }
	  ]
	}
*/

type WsAPIRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}
