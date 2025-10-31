package models

import (
	"encoding/json"

	"github.com/CrazyThursdayV50/goex/infra/utils"
)

type Request struct {
	ID     int64           `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

func (r *Request) SetID(id int64) {
	r.ID = id
}

func (r *Request) SetParams(params any) error {
	var err error
	r.Params, err = utils.JsonMarshalRaw(params)
	return err
}
