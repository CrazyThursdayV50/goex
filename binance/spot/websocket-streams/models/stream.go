package models

type StreamBase struct {
	Event  string `json:"e"`
	Symbol string `json:"s"`
	Time   int64  `json:"E"`
}
