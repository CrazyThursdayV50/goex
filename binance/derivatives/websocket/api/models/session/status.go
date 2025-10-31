package session

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"

type StatusParams = models.Request[any]

func Status() *StatusParams {
	return &StatusParams{
		Method: "session.status",
	}
}

type StatusResultData = LogonResultData
