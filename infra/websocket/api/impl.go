package api

import (
	"context"
	"errors"
	"time"

	"github.com/CrazyThursdayV50/goex/infra/iface"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
	"github.com/CrazyThursdayV50/pkgo/builtin"
	"github.com/CrazyThursdayV50/pkgo/builtin/wrap"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/google/uuid"
)

type API[Req iface.WebsocketRequest, Res iface.WebsocketResult] struct {
	logger        log.Logger
	client        *client.Client
	resultChan    chan Res
	resultMap     builtin.MapAPI[string, Res]
	resultTimeout time.Duration
	signerFunc    iface.SignerFunc
}

func (api *API[Req, Res]) Sign(requestData any) (map[string]any, error) {
	if api.signerFunc == nil {
		return nil, nil
	}
	return api.signerFunc(requestData)
}

func (api *API[Req, Res]) ReqId() string {
	return uuid.New().String()
}

func (api *API[Req, Res]) GetResult(ctx context.Context, id string) builtin.UnWrapper[Res] {
	ctx, cancel := context.WithTimeout(ctx, api.resultTimeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return wrap.Nil[Res]()

		default:
			if result := api.resultMap.Get(id); !result.IsNil() {
				return result
			}

			time.Sleep(time.Millisecond)
		}
	}
}

func (api *API[Req, Res]) Send(ctx context.Context, req Req) (builtin.UnWrapper[Res], error) {
	id := api.ReqId()
	req.SetID(id)

	data, err := req.MarshalBinary()
	if err != nil {
		return wrap.Nil[Res](), err
	}

	err = api.client.Send(data)
	if err != nil {
		return wrap.Nil[Res](), err
	}

	result := api.GetResult(ctx, id)
	if result.IsNil() {
		return wrap.Nil[Res](), errors.New("get result timeout")
	}

	res := result.Unwrap()
	if !res.IsOk() {
		return wrap.Nil[Res](), result.Unwrap()
	}

	return result, nil
}
