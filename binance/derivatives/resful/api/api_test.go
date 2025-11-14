package api

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/klines"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/trade"
	userdata "github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/useData"
	"github.com/CrazyThursdayV50/goex/binance/variables"
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
	variables.SetIsTest()

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

	apikey := os.Getenv("BN_APIKEY")
	secret := os.Getenv("BN_SECRET")
	api, err := New(&cfg, env.logger, nil, apikey, secret)
	if err != nil {
		t.Fatalf("create api failed: %v", err)
	}

	env.api = api

	return &env
}

func testAPI[Res any](name string, t *testing.T, testRequestFunc func(ctx context.Context) (Res, error), callback func(Res)) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	res, err := testRequestFunc(ctx)
	if err != nil {
		t.Fatalf("%s failed: %v", name, err)
	}
	t.Logf("%s: %+#v", name, res)

	if callback != nil {
		callback(res)
	}
}

func Test(t *testing.T) {
	env := Setup(t)
	api := env.api

	testAPI("Ping", t, api.Ping, nil)
	testAPI("ServerTime", t, api.ServerTime, nil)
	testAPI("ExchangeInfo", t, api.ExchangeInfo, nil)
	testAPI("Klines", t, func(ctx context.Context) (klines.Result, error) {
		return api.Klines(ctx, &klines.Params{
			Symbol:    env.symbol,
			Interval:  env.interval,
			StartTime: 0,
			EndTime:   0,
			Limit:     100,
		})
	}, nil)

	var currentMarginType userdata.MarginType
	var currentLeverage int64
	testAPI("SymbolConfig", t, func(ctx context.Context) (userdata.SymbolConfigResult, error) {
		return api.SymbolConfig(ctx, &userdata.SymbolConfigData{Symbol: env.symbol})
	}, func(res userdata.SymbolConfigResult) {
		for _, cfg := range res {
			if cfg.Symbol == env.symbol {
				currentMarginType = cfg.MarginType
				currentLeverage = cfg.Leverage
			}
		}
	})

	var marginType = userdata.MarginTypeCross
	if currentMarginType == marginType {
		marginType = userdata.MarginTypeIsolated
	}

	testAPI("ChangeMarginType", t, func(ctx context.Context) (*trade.MarginTypeResult, error) {
		return api.SetMarginType(ctx, &trade.MarginTypeParams{
			Symbol:     env.symbol,
			MarginType: marginType,
		})
	}, nil)
	testAPI("RestoreMarginType", t, func(ctx context.Context) (*trade.MarginTypeResult, error) {
		return api.SetMarginType(ctx, &trade.MarginTypeParams{
			Symbol:     env.symbol,
			MarginType: currentMarginType,
		})
	}, nil)

	var leverage int64 = 10
	if leverage == currentLeverage {
		leverage = 20
	}

	testAPI("ChangeOpenLeverage", t, func(ctx context.Context) (*trade.OpenLeverageResult, error) {
		return api.SetOpenLeverage(ctx, &trade.OpenLeverageParams{
			Symbol:   env.symbol,
			Leverage: leverage,
		})
	}, nil)
	testAPI("RestoreOpenLeverage", t, func(ctx context.Context) (*trade.OpenLeverageResult, error) {
		return api.SetOpenLeverage(ctx, &trade.OpenLeverageParams{
			Symbol:   env.symbol,
			Leverage: currentLeverage,
		})
	}, nil)
}
