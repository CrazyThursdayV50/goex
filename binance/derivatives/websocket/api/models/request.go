package models

import "github.com/CrazyThursdayV50/pkgo/json"

type Request struct {
	ID     string         `json:"id"`
	Method string         `json:"method"`
	Params map[string]any `json:"params,omitempty"`
}

func (p *Request) MarshalBinary() ([]byte, error) {
	return json.JSON().Marshal(p)
}

func (p *Request) SetID(id string) {
	p.ID = id
}
