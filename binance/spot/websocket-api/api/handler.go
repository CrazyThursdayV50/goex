package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-api/models"
	"github.com/CrazyThursdayV50/pkgo/builtin"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
)

func handler(responseMap builtin.MapAPI[string, *models.WsAPIResult]) func(context.Context, log.Logger, int, []byte, func(error)) (int, []byte) {

	return func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
		var result models.WsAPIResult
		err := json.JSON().Unmarshal(b, &result)
		if err != nil {
			l.Errorf("unmarshal failed: %v", err)
			return client.TextMessage, nil
		}

		responseMap.AddSoft(result.Id, &result)
		return client.TextMessage, nil
	}
}
