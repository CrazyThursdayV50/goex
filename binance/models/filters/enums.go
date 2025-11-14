package filters

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
