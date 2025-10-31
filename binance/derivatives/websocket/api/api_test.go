package api

import (
	"context"
	"os"
	"testing"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/apis"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/session"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/api/models/trade"
	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/infra/utils"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
	"github.com/shopspring/decimal"
)

var (
	logger log.Logger
	apikey string
	secret string
	api    *apis.API
	symbol string
)

func initLogger() {
	logger = sugar.New(sugar.DefaultConfig())
}

func initAPI(t *testing.T) {
	apikey = os.Getenv("BN_APIKEY")
	secret = os.Getenv("BN_SECRET")
	var err error
	api, err = apis.New(logger, apikey, secret)
	if err != nil {
		t.Fatalf("create api failed: %v", err)
	}

	err = api.Run(context.TODO())
	if err != nil {
		t.Fatalf("logon failed: %v", err)
	}
}

func Setup(t *testing.T) {
	symbol = "BTCUSDT"
	variables.SetIsTest()

	initLogger()
	initAPI(t)
}

func TestSession(t *testing.T) {
	Setup(t)
	ctx := context.TODO()

	r, err := api.Session().Logon(ctx, &session.LogonData{})
	if err != nil {
		t.Fatalf("logon failed: %v", err)
	}
	t.Logf("logon result: %v", r)

	status, err := api.Session().Status(ctx)
	if err != nil {
		t.Fatalf("get session status failed: %v", err)
	}

	t.Logf("session status: %+v", status)

	logout, err := api.Session().Logout(ctx)
	if err != nil {
		t.Fatalf("logon failed: %v", err)
	}

	t.Logf("session logout: %+v", logout)
}

// 单方向市价开多
func placeOrderOpenLongMarketDataSingle(symbol, quantity string) *trade.PlaceData {
	return &trade.PlaceData{
		Symbol:           symbol,
		OrderSide:        trade.SIDE_BUY,
		OrderType:        trade.TYPE_MARKET,
		Quantity:         utils.Ptr(quantity),
		NewOrderRespType: utils.Ptr(trade.NEW_ORDER_RESP_TYPE_RESULT),
	}
}

// 单方向多单市价止盈
func tpLongMarketSingle(symbol, stopPrice, quantity string) *trade.PlaceData {
	return &trade.PlaceData{
		Symbol:           symbol,
		OrderSide:        trade.SIDE_SELL,
		OrderType:        trade.TYPE_TAKE_PROFIT_MARKET,
		StopPrice:        utils.Ptr(stopPrice),
		ReduceOnly:       utils.Ptr(trade.REDUCE_ONLY_TRUE),
		Quantity:         utils.Ptr(quantity),
		NewOrderRespType: utils.Ptr(trade.NEW_ORDER_RESP_TYPE_RESULT),
	}
}

// 单方向多单市价止损
func slLongMarketSingle(symbol, stopPrice, quantity string) *trade.PlaceData {
	return &trade.PlaceData{
		Symbol:           symbol,
		OrderSide:        trade.SIDE_SELL,
		OrderType:        trade.TYPE_STOP_MARKET,
		StopPrice:        utils.Ptr(stopPrice),
		ReduceOnly:       utils.Ptr(trade.REDUCE_ONLY_TRUE),
		Quantity:         utils.Ptr(quantity),
		NewOrderRespType: utils.Ptr(trade.NEW_ORDER_RESP_TYPE_RESULT),
	}
}

// 单方向多单市价减仓
func reducePositionLongMarketDataSingle(symbol, quantity string) *trade.PlaceData {
	return &trade.PlaceData{
		Symbol:           symbol,
		OrderSide:        trade.SIDE_SELL,
		OrderType:        trade.TYPE_MARKET,
		ReduceOnly:       utils.Ptr(trade.REDUCE_ONLY_TRUE),
		Quantity:         utils.Ptr(quantity),
		NewOrderRespType: utils.Ptr(trade.NEW_ORDER_RESP_TYPE_RESULT),
	}
}

// 完全平仓
func TestTrade(t *testing.T) {
	Setup(t)

	var orderPrice decimal.Decimal
	t.Run("OpenLongMarket", func(t *testing.T) {
		var data trade.PlaceData
		data.SingleOpenLongMarket(symbol, "0.002")

		result, err := api.Trade().PlaceOrder(context.TODO(), &data)
		if err != nil {
			t.Fatalf("%s failed: %v", t.Name(), err)
		}

		orderPrice, _ = decimal.NewFromString(result.AvgPrice)
		t.Logf("result: %+v", result)
	})

	t.Run("TakeProfitLongMarket", func(t *testing.T) {
		stopPrice := orderPrice.Mul(decimal.NewFromFloat(1.1)).StringFixed(2)

		var data trade.PlaceData
		data.SingleLongTakeProfitMarket("BTCUSDT", "0.001", stopPrice)

		result, err := api.Trade().PlaceOrder(context.TODO(), &data)

		if err != nil {
			t.Fatalf("%s failed: %v", t.Name(), err)
		}
		t.Logf("result: %+v", result)
	})

	t.Run("StopLossLongMarket", func(t *testing.T) {
		stopPrice := orderPrice.Mul(decimal.NewFromFloat(0.9)).StringFixed(2)
		var data trade.PlaceData
		data.SingleLongStopLossMarket("BTCUSDT", "0.001", stopPrice)
		result, err := api.Trade().PlaceOrder(context.TODO(), &data)

		if err != nil {
			t.Fatalf("%s failed: %v", t.Name(), err)
		}
		t.Logf("result: %+v", result)
	})

	t.Run("ReduceLongMarket", func(t *testing.T) {
		var data trade.PlaceData
		data.SingleReduceLongMarket("BTCUSDT", "0.001")

		result, err := api.Trade().PlaceOrder(context.TODO(), &data)
		if err != nil {
			t.Fatalf("%s failed: %v", t.Name(), err)
		}
		t.Logf("result: %+v", result)
	})
}
