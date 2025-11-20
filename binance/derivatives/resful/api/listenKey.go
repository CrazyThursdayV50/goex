package api

import (
	"context"

	listenkey "github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/listenKey"
)

const (
	listenKeyPath = "/fapi/v1/listenKey"
)

func (api *API) NewListenKey(ctx context.Context) (result *listenkey.PostResult, err error) {
	err = api.userstream().post(ctx, listenKeyPath, nil, &result)
	return
}

func (api *API) ExtendListenKey(ctx context.Context) (result *listenkey.PutResult, err error) {
	err = api.userstream().put(ctx, listenKeyPath, nil, &result)
	return
}

func (api *API) RemoveListenKey(ctx context.Context) (result *listenkey.DeleteResult, err error) {
	err = api.userstream().delete(ctx, listenKeyPath, nil, &result)
	return
}
