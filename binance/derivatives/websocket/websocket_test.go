package websocket

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/apis"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models/account"
	listenkey "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models/listenKey"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models/order"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/streams"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
)

var (
	logger     log.Logger
	apikey     string
	secret     string
	wsapi      *apis.API
	userStream *streams.Stream
	symbol     string
	listenKey  string
)

func initLogger() {
	logger = sugar.New(sugar.DefaultConfig())
}

func initAPI(t *testing.T) {
	apikey = os.Getenv("BN_APIKEY")
	secret = os.Getenv("BN_SECRET")
	var err error
	wsapi, err = apis.New(logger, apikey, secret)
	if err != nil {
		t.Fatalf("create api failed: %v", err)
	}

	err = wsapi.Run(context.TODO())
	if err != nil {
		t.Fatalf("logon failed: %v", err)
	}

	result, err := wsapi.NewListenKey(t.Context())
	if err != nil {
		t.Fatalf("NewListenKey failed: %v", err)
	}

	listenKey = result.ListenKey
}

func defaultHandler[Event any](t *testing.T, name string) func(Event, error) {
	return func(e Event, err error) {
		if err != nil {
			t.Fatalf("%s failed: %v", name, err)
			return
		}
		t.Logf("%s: %v", name, e)
	}
}

func initStream(t *testing.T) {
	userStream = streams.New(logger, listenKey)
	userStream.HandleListenKeyExpired(defaultHandler[*listenkey.Event](t, listenkey.EventName))
	userStream.HandleAccountUpdate(defaultHandler[*account.Event](t, account.EventName))
	userStream.HandleOrderUpdate(defaultHandler[*order.Event](t, order.EventName))

	userStream.Run(t.Context())
}

func Setup(t *testing.T) {
	symbol = "BTCUSDT"
	variables.SetIsTest()

	initLogger()
	initAPI(t)
	initStream(t)
}

func Test(t *testing.T) {
	Setup(t)

	time.Sleep(time.Hour)
}
