package session

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"
)

type LogonData struct {
	models.Sign
}

// implement Mapper
func (d *LogonData) Map() map[string]string {
	m := d.Sign.Map()
	return m
}

type LogonRequest = models.Request[*LogonData]

func Logon() *LogonRequest {
	return &LogonRequest{
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
