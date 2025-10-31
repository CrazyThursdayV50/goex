package session

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"
)

type LogonData struct {
	models.Sign
}

// implement Mapper
func (p *LogonData) Map() map[string]string {
	m := p.Sign.Map()
	return m
}

type LogonParams = models.Request[*LogonData]

func Logon() *LogonParams {
	return &LogonParams{
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
