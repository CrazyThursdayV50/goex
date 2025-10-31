package client

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
)

type Client struct {
	logger   log.Logger
	WsClient *client.Client
}

func (c *Client) OnConnect(f func() (int, []byte)) {
	c.WsClient.UpdateOptions(client.WithSendOnConnect(f))
}

func (c *Client) Stop()                         { c.WsClient.Stop() }
func (c *Client) Run(ctx context.Context) error { return c.WsClient.Run(ctx) }

var PingMessage = client.PingMessage
var TextMessage = client.TextMessage
var BinaryMessage = client.BinaryMessage

type MessageHandler = client.MessageHandler

func NewClient(
	logger log.Logger,
	url, proxy string,
	handler client.MessageHandler,
	pingLoop client.PingLoop,
) *Client {
	logger.Debugf("websocket connecting to %s ...", url)
	wsclient := client.New(
		client.WithProxy(proxy),
		client.WithLogger(logger),
		client.WithURL(url),
		client.WithDefaultCompress(true),
		client.WithPingHandler(variables.WriteControlTimeout(), func(string) error {
			logger.Debugf("Recv: PING")
			return nil
		}),
		client.WithPongHandler(variables.WriteControlTimeout(), func(string) error {
			logger.Debugf("Recv: PONG")
			return nil
		}),
		client.WithPingLoop(pingLoop),
		client.WithMessageHandler(func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			switch i {
			case client.TextMessage:
				logger.Debugf("Recv: TEXT, %s", b)
				return handler(ctx, l, i, b, f)

			default:
				return client.BinaryMessage, nil
			}
		}),
	)

	return &Client{logger: logger, WsClient: wsclient}
}

func (c *Client) Ping(data []byte) error {
	return c.WsClient.Ping(data)
}

func (c *Client) Pong(data []byte) error {
	return c.WsClient.Pong(data)
}

func (c *Client) Send(data []byte) error {
	c.logger.Debugf("Send: TEXT, %s", data)
	return c.WsClient.Send(data)
}
