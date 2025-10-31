package session

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"

type LogoutParams = models.Request[any]

func Logout() *LogoutParams { return &LogoutParams{Method: "session.logout"} }

type LogoutResultData = LogonResultData
