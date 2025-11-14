package api

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/apis"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type API struct{}

func New() API { return API{} }

func (a API) Public(logger log.Logger) *apis.API {
	api, _ := apis.New(logger, "", "")
	return api
}

func (a API) Private(logger log.Logger, apiKey, secretKey string) (*apis.API, error) {
	return apis.New(logger, apiKey, secretKey)
}
