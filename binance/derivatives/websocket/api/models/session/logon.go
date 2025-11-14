package session

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"

type LogonData struct {
	RecvWindow int64 `json:"recvWindow,omitempty"`
}

func Logon() *models.Request {
	return &models.Request{
		Method: "session.logon",
	}
}

type LogonResultData struct {
	APIKEY           string `json:"apiKey"`
	AuthorizedSince  int64  `json:"authorizedSince"`
	ConnectedSince   int64  `json:"connectedSince"`
	ReturnRateLimits bool   `json:"returnRateLimits"`
	ServerTime       int64  `json:"serverTime"`
}
