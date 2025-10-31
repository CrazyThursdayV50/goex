package iface

import (
	"context"
	"encoding"
	"encoding/json"

	"github.com/CrazyThursdayV50/pkgo/builtin"
)

type WebsocketRequest interface {
	SetID(string)
	encoding.BinaryMarshaler
}

type WebsocketResult interface {
	IsOk() bool
	Message() string
	error

	GetID() string
	Data() json.RawMessage
}

type WebsocketAPI[Req WebsocketRequest, Res WebsocketResult] interface {
	Sign(any) (map[string]any, error)
	ReqId() string

	Send(ctx context.Context, req Req) (builtin.UnWrapper[Res], error)
	GetResult(ctx context.Context, id string) builtin.UnWrapper[Res]
}
