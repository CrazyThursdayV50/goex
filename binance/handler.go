package binance

import (
	"context"

	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
)

var handlePing client.MessageHandler = func(_ context.Context, _ log.Logger, typ int, data []byte, _ func(error)) (int, []byte) {
	if typ == client.PingMessage {
		return client.PongMessage, data
	}

	return client.BinaryMessage, nil
}
