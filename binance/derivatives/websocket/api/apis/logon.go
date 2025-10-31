package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/session"
)

func (api *API) Logon(ctx context.Context) (result *session.LogonResultData, err error) {
	var data session.LogonData
	api.Sign(&data)
	params := session.Logon()
	params.Params = &data
	err = request(ctx, api, params, &result)
	return
}
