package api

import (
	"context"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/infra/iface"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
	gmap "github.com/CrazyThursdayV50/pkgo/builtin/map"
	"github.com/CrazyThursdayV50/pkgo/cron"
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/worker"
	"github.com/gorilla/websocket"
)

func New[Req iface.WebsocketRequest, Res iface.WebsocketResult](
	logger log.Logger,
	endpoint string,
	signerFunc iface.SignerFunc,
	handler func(chan Res) client.MessageHandler,
) *API[Req, Res] {
	resultMap := gmap.Make[string, Res](0)

	pingLoop := func(done <-chan struct{}, conn *websocket.Conn) {
		ctx, cancel := context.WithCancel(context.TODO())
		goo.Go(func() {
			<-done
			cancel()
		})

		cron.New(
			cron.WithJob(func() {
				logger.Debugf("[%s]PING sent", endpoint)
				conn.WriteControl(client.PingMessage, nil, time.Now().Add(variables.WriteControlTimeout()))
			}, time.Second*10),
			cron.WithLogger(logger),
			cron.WithRunAfterStart(time.Second),
			cron.WithWaitAfterRun(false),
		).Run(ctx)
	}

	ch := make(chan Res, 100)
	c := client.NewClient(
		logger,
		endpoint,
		variables.GetProxy(),
		handler(ch),
		pingLoop,
	)

	api := API[Req, Res]{
		client:        c,
		resultChan:    ch,
		resultMap:     resultMap,
		logger:        logger,
		resultTimeout: time.Second * 10,
		signerFunc:    signerFunc,
	}

	return &api
}

func (api *API[Req, Res]) Run(ctx context.Context) error {
	err := api.client.Run(ctx)
	if err != nil {
		return err
	}

	worker, _ := worker.New("WsResultReceive", func(r Res) {
		api.resultMap.AddSoft(r.GetID(), r)
	})

	worker.WithLogger(api.logger)
	worker.WithGraceful(true)
	worker.WithTrigger(api.resultChan)
	worker.Run(ctx)
	return nil
}

func (api *API[Req, Res]) Stop() {
	api.client.Stop()
}

func (api *API[Req, Res]) SetGetResultTimeout(timeout time.Duration) {
	api.resultTimeout = timeout
}
