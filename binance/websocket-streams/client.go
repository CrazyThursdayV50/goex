package websocketstreams

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
)

type Client struct {
	WsClient *client.Client
}

func (c *Client) Stop()      { c.WsClient.Stop() }
func (c *Client) Run() error { return c.WsClient.Run() }

func NewClient(ctx context.Context, logger log.Logger, url string, handler client.MessageHandler) *Client {
	logger.Debugf("ws connect to %s", url)
	wsclient := client.New(
		client.WithContext(ctx),
		client.WithProxy(variables.GetProxy()),
		client.WithLogger(logger),
		client.WithURL(url),
		client.WithDefaultCompress(true),
		client.WithPingHandler(variables.WriteControlTimeout(), nil),
		client.WithPongHandler(variables.WriteControlTimeout(), func(string) error {
			logger.Debugf("PONG recv")
			return nil
		}),
		client.WithMessageHandler(func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			switch i {
			case client.TextMessage:
				return handler(ctx, l, i, b, f)

			default:
				return client.BinaryMessage, nil
			}
		}),
	)

	return &Client{WsClient: wsclient}
}

func (c *Client) Ping(data []byte) error {
	return c.WsClient.Ping(data)
}

func (c *Client) Pong(data []byte) error {
	return c.WsClient.Pong(data)
}

func (c *Client) Send(data []byte) error {
	return c.WsClient.Send(data)
}
