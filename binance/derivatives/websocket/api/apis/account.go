package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/account"
)

type Account struct {
	api *API
}

func (api *API) Account() *Account {
	return &Account{
		api: api,
	}
}

func (a *Account) Status(ctx context.Context, data *account.StatusData) (result *account.StatusResultData, err error) {
	req := account.Status()
	req.Params, err = a.api.Sign(data)
	if err != nil {
		return
	}

	err = request(ctx, a.api.api, req, &result)
	return
}
