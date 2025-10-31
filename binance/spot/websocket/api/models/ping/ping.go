package ping

import (
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models"
)

// Params Ping请求参数
type ParamsData struct{}

func NewParams() *models.WsAPIRequest {
	return &models.WsAPIRequest{
		Method: "ping",
	}
}

type ResultData struct{}
