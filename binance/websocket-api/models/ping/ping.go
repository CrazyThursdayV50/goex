package ping

import (
	"fmt"

	"github.com/CrazyThursdayV50/goex/binance/websocket-api/models"
	"github.com/CrazyThursdayV50/pkgo/json"
)

// Params Ping请求参数
type Params = models.WsAPIParams[any]

func NewParams() *Params {
	return &Params{
		Method: "ping",
	}
}

type ResultData struct{}

func (p *ResultData) UnmarshalBinary(data []byte) error {
	if p == nil {
		return fmt.Errorf("nil receiver")
	}
	return json.JSON().Unmarshal(data, p)
}
