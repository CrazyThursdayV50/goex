package apis

import (
	"context"

	listenkey "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/listenKey"
)

func (api *API) NewListenKey(ctx context.Context) (result *listenkey.ResultData, err error) {
	req := listenkey.Start()
	req.Params = map[string]any{
		"apiKey": api.apikey,
	}

	err = request(ctx, api.api, req, &result)
	return
}

func (api *API) ExtendListenKey(ctx context.Context) (result *listenkey.ResultData, err error) {
	req := listenkey.Ping()
	req.Params = map[string]any{
		"apiKey": api.apikey,
	}

	err = request(ctx, api.api, req, &result)
	return
}

func (api *API) RemoveListenKey(ctx context.Context) (result *listenkey.StopResultData, err error) {
	req := listenkey.Start()
	req.Params = map[string]any{
		"apiKey": api.apikey,
	}

	err = request(ctx, api.api, req, &result)
	return
}
