package models

type BaseEvent struct {
	EventName string `json:"e"`
	EventTime int64  `json:"E"`
	MatchTime int64  `json:"T,omitempty"` // 撮合时间
}
