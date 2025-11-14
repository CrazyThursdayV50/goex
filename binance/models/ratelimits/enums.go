package ratelimits

type RateLimitType string

const (
	// 单位时间请求权重之和上限
	REQUESTS_WEIGHT RateLimitType = "REQUESTS_WEIGHT"
	// 单位时间下单(撤单)次数上限
	ORDERS RateLimitType = "ORDERS"
	// 单位时间请求次数上限
	RAW_REQUESTS RateLimitType = "RAW_REQUESTS"
)
