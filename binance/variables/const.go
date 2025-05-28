package variables

import "time"

var istest = false
var proxy = ""

func SetIsTest()          { istest = true }
func SetProxy(url string) { proxy = url }
func GetProxy() string    { return proxy }

var wsapiReadMessageTimeout = time.Second * 10

func SetWsAPIReadMessageTimeout(duration time.Duration) {
	wsapiReadMessageTimeout = duration
}

var writeControlTimeout = time.Second

func SetWriteControlTimeout(timeout time.Duration) { writeControlTimeout = timeout }
func WriteControlTimeout() time.Duration           { return writeControlTimeout }

func WsAPIReadmessageTimeout() time.Duration { return wsapiReadMessageTimeout }

// 本篇所列出的所有wss接口的baseurl为: wss://stream.binance.com:9443 或者 wss://stream.binance.com:443
const WS_STREAM_BASE_URL = "wss://stream.binance.com:443"
const WS_STREAM_TESAT_BASE_URL = "wss://stream.testnet.binance.vision/ws"

/*
本篇所列出的 wss 接口的 base URL：wss://ws-api.binance.com:443/ws-api/v3
如果使用标准443端口时遇到问题，可以使用替代端口9443。
现货测试网的 base URL 是 wss://ws-api.testnet.binance.vision/ws-api/v3。
*/
const WS_API_BASE_URL = "wss://ws-api.binance.com:443/ws-api/v3"
const WS_API_TEST_BASE_URL = "wss://ws-api.testnet.binance.vision/ws-api/v3"

// const BASE_URL = "wss://stream.binance.com:9443"
const STREAM_URL = "/ws/%s"
const COMBINED_STREAM_URL = "/stream?streams=%s"

// Stream Names: <symbol>@depth<levels> OR <symbol>@depth<levels>@100ms
const PARTIAL_BOOK_DEPTH = "%s@depth%d"
const PARTIAL_BOOK_DEPTH_100ms = "%s@depth%d@100ms"
const INDIVIDUAL_BOOK_TICKER = "%s@bookTicker"

func streamBaseUrl() string {
	if istest {
		return WS_STREAM_TESAT_BASE_URL
	}

	return WS_STREAM_BASE_URL
}

func StreamURL() string {
	return streamBaseUrl() + STREAM_URL
}

func CombinedStreamURL() string {
	return streamBaseUrl() + COMBINED_STREAM_URL
}

func WsAPIURL() string {
	if istest {
		return WS_API_TEST_BASE_URL
	}
	return WS_API_BASE_URL
}
