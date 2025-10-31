package streams

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/subscribe"
)

// Subscribe
// dataHandler can be created by CreateBytesHandler
func (s *Stream) Subscribe(ctx context.Context, params subscribe.RequestParams) error {
	req := subscribe.NewSubscribe()
	req.Params = params
	return request(s, req)
}

func (s *Stream) Unsubscribe(ctx context.Context, params subscribe.RequestParams) error {
	req := subscribe.NewUnsubscribe()
	req.Params = params
	return request(s, req)
}
