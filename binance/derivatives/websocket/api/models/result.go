package models

import (
	"context"
	j "encoding/json"
	"fmt"

	"github.com/CrazyThursdayV50/pkgo/builtin"
)

type WsRequest interface {
	Do(ctx context.Context) (builtin.UnWrapper[j.Unmarshaler], error)
}

type ResultError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ResultError) String() string {
	if e == nil {
		return "nil"
	}

	return fmt.Sprintf("[%d]%s", e.Code, e.Msg)
}

type Result struct {
	ID         string       `json:"id"`
	Status     int          `json:"status"`
	Result     j.RawMessage `json:"result"`
	RateLimits []RateLimit  `json:"rateLimits"`
	Error      *ResultError `json:"error,omitempty"`
}

func (r *Result) String() string {
	return fmt.Sprintf("[%s] - %d - %s", r.ID, r.Status, r.Result)
}
