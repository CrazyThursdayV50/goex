package models

type Request[Params any] struct {
	ID     string `json:"id"`
	Method string `json:"method"`
	Params Params `json:"params,omitempty"`
}
