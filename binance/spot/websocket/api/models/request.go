package models

import (
	"github.com/CrazyThursdayV50/pkgo/json"
)

// type WsRequest interface {
// 	Do(ctx context.Context) (builtin.UnWrapper[j.Unmarshaler], error)
// }

// WsAPIRequest WebSocket API 请求参数基础结构
type WsAPIRequest struct {
	ID     string         `json:"id"`
	Method string         `json:"method"`
	Params map[string]any `json:"params,omitempty"`
}

func (p *WsAPIRequest) MarshalBinary() ([]byte, error) {
	return json.JSON().Marshal(p)
}

// func (p *WsAPIParams[T]) BinaryMarshal() ([]byte, error) {
// 	return json.JSON().Marshal(p)
// }

// func (p *WsAPIParams[T]) BinaryUnmarshal(data []byte) error {
// 	if p == nil {
// 		return errors.New("nil receiver")
// 	}

// 	return json.JSON().Unmarshal(data, p)
// }

func (p *WsAPIRequest) SetID(id string) {
	p.ID = id
}
