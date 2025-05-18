package websocketapi

import (
	"context"
	"goex/binance/models"
)

func (api *API) MarketOrder(ctx context.Context, symbol, marketQuantity string, isBuy bool) (*models.WsOrderResultData, error) {
	order := models.WsOrderParamsData{
		Type: models.MARKET,
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
