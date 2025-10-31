package models

import "encoding/json"

type Response struct {
	ID     int64           `json:"id"`
	Result json.RawMessage `json:"result"`
}
