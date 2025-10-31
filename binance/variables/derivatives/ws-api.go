package derivatives

import (
	"time"

	"github.com/CrazyThursdayV50/goex/binance/variables"
)

const urlBaseWebsocketAPI = "wss://ws-fapi.binance.com"
const urlBaseWebsocketAPITest = "wss://testnet.binancefuture.com"
const pathWebsocketAPI = "/ws-fapi/v1"

type api struct{}

func WsAPI() api { return api{} }

func apiBaseURL() string {
	if variables.IsTest() {
		return urlBaseWebsocketAPITest
	}
	return urlBaseWebsocketAPI
}

func (api) Endpoint() string { return apiBaseURL() + pathWebsocketAPI }

// websocket api 接收数据超市时间
var readMessageTimeoutWebSocketAPI = time.Second * 10

func (api) ReadMessageTimeout() time.Duration { return readMessageTimeoutWebSocketAPI }
func (api) SetReadMessageTimeout(duration time.Duration) {
	readMessageTimeoutWebSocketAPI = duration
}
