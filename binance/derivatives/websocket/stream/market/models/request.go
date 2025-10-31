package models

type Request[Params any] struct {
	ID     int64  `json:"id"`
	Method string `json:"method"`
	Params Params `json:"params"`
}

func (r *Request[Params]) SetID(id int64) {
	r.ID = id
}
