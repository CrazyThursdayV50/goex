package websocketapi

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models"
	"github.com/CrazyThursdayV50/pkgo/builtin"
)

func (api *API) MarketOrder(ctx context.Context, symbol, marketQuantity string, isBuy bool) (builtin.UnWrapper[*models.WsOrderResultData], error) {
	order := models.WsOrderParamsData{
		Type:   models.MARKET,
		Symbol: symbol,
	}

	switch isBuy {
	case true:
		order.Side = models.BUY
		order.QuoteOrderQty = &marketQuantity

	default:
		order.Side = models.SELL
		order.Quantity = &marketQuantity
	}

	return api.Order(ctx, &order)
}
