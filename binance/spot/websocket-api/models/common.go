package models

import (
	"context"
	"encoding"
	"errors"
	"fmt"
	"time"

	"github.com/CrazyThursdayV50/pkgo/builtin"
	"github.com/CrazyThursdayV50/pkgo/json"
)

type ResultData interface {
	UnmarlshalJSON([]byte) error
}

type WsRequest interface {
	Do(ctx context.Context) (builtin.UnWrapper[ResultData], error)
}

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

func (p *WsAPIParams[T]) SetId(id string) { p.Id = id }

// WsAPIResult WebSocket API 响应结果基础结构
type WsAPIResult struct {
	Id         string                     `json:"id"`
	Status     int                        `json:"status"`
	Result     json.RawMessage            `json:"result"`
	RateLimits []*RateLimit               `json:"rateLimits"`
	Err        *WsAPIResultError          `json:"error,omitempty"`
	InnerData  encoding.BinaryUnmarshaler `json:"-"`
}

/*
	"error": {
	  "code": -2010,
	  "msg": "Account has insufficient balance for requested action."
	},
*/

type WsAPIResultError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func (e *WsAPIResultError) String() string {
	if e == nil {
		return "nil"
	}

	return fmt.Sprintf("[%d]%s", e.Code, e.Msg)
}

func (r *WsAPIResult) String() string {
	return fmt.Sprintf("[%s] - %d - %s", r.Id, r.Status, r.Result)
}

func (r *WsAPIResult) IsOk() bool { return r.Status == 200 }

func (r *WsAPIResult) Error() string {
	return fmt.Sprintf("status: %d, error: %s", r.Status, r.Err)
}

func (r *WsAPIResult) UnmarshalData() error {
	return r.InnerData.UnmarshalBinary(r.Result)
}

// RateLimit API速率限制信息
type RateLimit struct {
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
