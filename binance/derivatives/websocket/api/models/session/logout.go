package session

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"

type LogoutData struct{}

func Logout() *models.Request { return &models.Request{Method: "session.logout"} }

type LogoutResultData = LogonResultData
