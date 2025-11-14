package models

import (
	"encoding"
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/models/ratelimits"
	"github.com/CrazyThursdayV50/pkgo/json"
)

// WsAPIResult WebSocket API 响应结果基础结构
type WsAPIResult struct {
	Id         string                     `json:"id"`
	Status     int                        `json:"status"`
	Result     json.RawMessage            `json:"result"`
	RateLimits []ratelimits.RateLimit     `json:"rateLimits"`
	Err        *WsAPIResultError          `json:"error,omitempty"`
	InnerData  encoding.BinaryUnmarshaler `json:"-"`
}

func (r *WsAPIResult) GetID() string { return r.Id }

func (r *WsAPIResult) Data() json.RawMessage { return r.Result }

func (r *WsAPIResult) Message() string {
	if r == nil {
		return "nil result"
	}

	if r.Err == nil {
		return ""
	}

	return r.Err.Msg
}

func (r *WsAPIResult) String() string {
	return fmt.Sprintf("[%s] - %d - %s", r.Id, r.Status, r.Result)
}

func (r *WsAPIResult) IsOk() bool { return r != nil && r.Status == 200 }

func (r *WsAPIResult) Error() string {
	if r == nil {
		return ""
	}

	return r.Err.String()
}

func (r *WsAPIResult) UnmarshalData() error {
	return r.InnerData.UnmarshalBinary(r.Result)
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

func (e *WsAPIResultError) Error() string {
	return e.String()
}
