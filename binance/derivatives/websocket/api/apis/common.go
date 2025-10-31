package apis

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"
	"github.com/CrazyThursdayV50/goex/infra/iface"
	"github.com/CrazyThursdayV50/pkgo/json"
)

func request[Data any](
	ctx context.Context,
	api iface.WebsocketAPI[*models.Request, *models.Result],
	params *models.Request,
	result *Data,
) error {
	res, err := api.Send(ctx, params)
	if err != nil {
		return err
	}

	err = json.JSON().Unmarshal(res.Unwrap().Data(), result)
	if err != nil {
		return err
	}

	return nil
}
