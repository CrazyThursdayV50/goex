package models

import (
	"encoding/json"
	j "encoding/json"
	"fmt"
)

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
	Err        *ResultError `json:"error,omitempty"`
}

func (r *Result) String() string {
	return fmt.Sprintf("[%s] - %d - %s", r.ID, r.Status, r.Result)
}

func (r *Result) GetID() string { return r.ID }
func (r *Result) IsOk() bool    { return r != nil && r.Status == 200 }
func (r *Result) Message() string {
	if r == nil {
		return "nil result"
	}

	if r.Err == nil {
		return ""
	}

	return r.Err.Msg
}

func (r *Result) Data() json.RawMessage { return r.Result }

func (r *Result) Error() string {
	if r == nil {
		return ""
	}

	return r.Err.String()
}
