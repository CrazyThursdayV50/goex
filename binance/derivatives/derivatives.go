package derivatives

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/api"
	wsapi "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api"
	wsstream "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/request/resty"
	"github.com/CrazyThursdayV50/pkgo/trace"
)

type DerivativesEntry struct{}

func (DerivativesEntry) WebSocketAPI() wsapi.API {
	return wsapi.New()
}

func (DerivativesEntry) WebSocketStream() wsstream.Stream {
	return wsstream.New()
}

func (DerivativesEntry) Restful(
	cfg *resty.Config,
	logger log.Logger,
	tracer trace.Tracer,
	apiKey, secret string,
) (*api.API, error) {
	return api.New(cfg, logger, tracer, apiKey, secret)
}
