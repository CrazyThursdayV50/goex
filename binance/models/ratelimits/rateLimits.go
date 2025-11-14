package ratelimits

// RateLimit API速率限制信息
type RateLimit struct {
	RateLimitType RateLimitType `json:"rateLimitType"`
	Interval      string        `json:"interval"`
	IntervalNum   int           `json:"intervalNum"`
	Limit         int           `json:"limit"`
	Count         int           `json:"count"`
}
