package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models/exchangeinfo"
	"github.com/CrazyThursdayV50/goex/infra/utils"
)

func (api *API) ExchangeInfo(ctx context.Context, data *exchangeinfo.ParamsData) (result *exchangeinfo.ResultData, err error) {
	req := exchangeinfo.NewRequest()
	req.Params, _ = utils.MapAny(data)

	err = request(ctx, api.api, req, &result)
	return
}
