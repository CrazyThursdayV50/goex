package api

import (
	"context"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/klines"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
	"github.com/CrazyThursdayV50/pkgo/request/resty"
)

type TestEnv struct {
	symbol   string
	interval string

	logger log.Logger
	api    *API
}

func Setup(t *testing.T) *TestEnv {
	var env TestEnv
	env.symbol = "BTCUSDT"
	env.interval = "1m"

	env.logger = sugar.New(sugar.DefaultConfig())

	var cfg resty.Config
	cfg.Debug = false
	cfg.EnableTrace = false
	cfg.RetryCount = 3
	cfg.RetryMaxWaitTime = time.Second * 2
	cfg.RetryWaitTime = time.Second
	cfg.Timeout = time.Minute
	cfg.EnableLog = true

	env.api = New(&cfg, env.logger, nil)
	return &env
}

func testAPI[Res any](name string, t *testing.T, testRequestFunc func(ctx context.Context) (Res, error)) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	res, err := testRequestFunc(ctx)
	if err != nil {
		t.Fatalf("%s failed: %v", name, err)
	}
	t.Logf("%s: %v", name, res)
}

func Test(t *testing.T) {
	env := Setup(t)
	api := env.api

	testAPI("Ping", t, api.Ping)
	testAPI("ServerTime", t, api.ServerTime)
	testAPI("Klines", t, func(ctx context.Context) (klines.Result, error) {
		return api.Klines(ctx, &klines.Params{
			Symbol:    env.symbol,
			Interval:  env.interval,
			StartTime: 0,
			EndTime:   0,
			Limit:     100,
		})
	})
}
