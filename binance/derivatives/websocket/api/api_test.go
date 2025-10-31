package api

import (
	"context"
	"os"
	"testing"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/apis"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
)

var (
	logger log.Logger
	apikey string
	secret string
)

func initLogger(t *testing.T) {
	logger = sugar.New(sugar.DefaultConfig())
}

func initAPI(t *testing.T) {
	apikey = os.Getenv("BN_APIKEY")
	secret = os.Getenv("BN_SECRET")
}

func Setup(t *testing.T) {
	initLogger(t)
	initAPI(t)
}

func TestLogon(t *testing.T) {
	Setup(t)
	ctx := context.TODO()

	api := apis.New(logger, apikey, secret)
	err := api.Run(ctx)
	if err != nil {
		t.Fatalf("logon failed: %v", err)
	}
}
