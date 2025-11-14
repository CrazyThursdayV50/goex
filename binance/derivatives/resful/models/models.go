package models

import "fmt"

type ResultError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func (r *ResultError) Error() string {
	if r == nil {
		return "nil"
	}

	return fmt.Sprintf("[%d]%s", r.Code, r.Msg)
}
