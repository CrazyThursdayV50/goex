package api

import (
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/apis"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type API struct{}

func New() API {
	return API{}
}

func (API) Private(logger log.Logger, apiKey, secretKey string) (*apis.API, error) {
	return apis.New(logger, apiKey, secretKey)
}

func (API) Public(logger log.Logger) *apis.API {
	apis, _ := apis.New(logger, "", "")
	return apis
}
