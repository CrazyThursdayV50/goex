package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/CrazyThursdayV50/pkgo/json"
)

// WsAPIParams WebSocket API 请求参数基础结构
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

// WsAPIResult WebSocket API 响应结果基础结构
type WsAPIResult struct {
	Id         string            `json:"id"`
	Status     int               `json:"status"`
	Result     json.RawMessage   `json:"result"`
	RateLimits []*WsAPIRateLimit `json:"rateLimits"`
}

func (r *WsAPIResult) String() string {
	return fmt.Sprintf("[%s] - %d - %s", r.Id, r.Status, r.Result)
}

// WsAPIRateLimit API速率限制信息
type WsAPIRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

// Sign 签名基础结构
type Sign struct {
	ApiKey    string `json:"apiKey"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

func (p *Sign) SetTimestamp() {
	p.Timestamp = time.Now().UnixMilli()
}

func (p *Sign) SetApiKey(apikey string) {
	p.ApiKey = apikey
}

func (p *Sign) SetSignature(signature string) {
	p.Signature = signature
}

func (p *Sign) Map() map[string]string {
	return map[string]string{
		"apiKey":    p.ApiKey,
		"timestamp": fmt.Sprintf("%d", p.Timestamp),
	}
} 