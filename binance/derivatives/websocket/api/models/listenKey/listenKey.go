package listenkey

import "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models"

type Data struct{}

func Start() *models.Request {
	return &models.Request{
		Method: "userDataStream.start",
	}
}
func Ping() *models.Request {
	return &models.Request{
		Method: "userDataStream.ping",
	}
}
func Stop() *models.Request {
	return &models.Request{
		Method: "userDataStream.stop",
	}
}

type ResultData struct {
	ListenKey string `json:"listenKey"`
}

type StopResultData struct{}
