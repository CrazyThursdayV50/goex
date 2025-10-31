package spot

import (
	"time"

	"github.com/CrazyThursdayV50/goex/binance/variables"
)

/*
本篇所列出的 wss 接口的 base URL：wss://ws-api.binance.com:443/ws-api/v3
如果使用标准443端口时遇到问题，可以使用替代端口9443。
现货测试网的 base URL 是 wss://ws-api.testnet.binance.vision/ws-api/v3。
*/

const urlBaseWebsocketAPI = "wss://ws-api.binance.com:443"
const urlBaseWebsocketAPITest = "wss://ws-api.testnet.binance.vision"
const pathWebsocketAPI = "/ws-api/v3"

type api struct{}

func WsAPI() api { return api{} }

func apiBaseURL() string {
	if variables.IsTest() {
		return urlBaseWebsocketAPITest
	}
	return urlBaseWebsocketAPI
}

func (api) Endpoint() string {
	return apiBaseURL() + pathWebsocketAPI
}

// websocket 接收数据超市时间
var wsapiReadMessageTimeout = time.Second * 10

func (api) ReadMessageTimeout() time.Duration { return wsapiReadMessageTimeout }
func (api) SetReadMessageTimeout(duration time.Duration) {
	wsapiReadMessageTimeout = duration
}
