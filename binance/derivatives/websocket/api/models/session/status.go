package session

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"

type StatusData struct{}

func Status() *models.Request {
	return &models.Request{
		Method: "session.status",
	}
}

type StatusResultData = LogonResultData
