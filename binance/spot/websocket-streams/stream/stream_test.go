package stream

import (
	"context"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-streams/models"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket-streams/models/klines"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
)

func TestStreamCreation(t *testing.T) {
	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())

	symbol := "BTCUSDT"

	// 测试创建流客户端但不连接
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("创建流客户端时发生panic: %v", r)
		}
	}()

	handler := func(event *models.PartialDepthData) {
		logger.Infof("depth: %s", event.String())
	}

	// 创建客户端但立即停止，只测试创建过程
	client := PartialBookDepth5Stream(ctx, logger, symbol, handler)
	if client == nil {
		t.Errorf("创建流客户端失败")
		return
	}

	// 立即停止，不等待连接
	client.Stop()
	t.Log("流客户端创建测试通过")
}

func TestStream(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())
	symbol := "BTCUSDT"
	handler := func(event *models.PartialDepthData) {
		logger.Infof("depth: %s", event.String())
	}

	client := PartialBookDepth5Stream(ctx, logger, symbol, handler)
	time.Sleep(time.Second * 2) // 缩短测试时间
	client.Stop()
}

func TestCombinedStream(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())
	symbols := []string{"SOLBTC", "BNBBTC", "BNBUSDT"}
	handler := func(event *models.PartialDepthCombinedData) {
		logger.Infof("depth: %s", event.String())
	}

	client := PartialBookDepth5CombinedStream(ctx, logger, symbols, handler)
	time.Sleep(time.Second * 2) // 缩短测试时间
	client.Stop()
}

func TestIndividualSymbolBookTicker(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())
	symbols := []string{"SOLBTC", "BNBBTC", "BNBUSDT"}
	handler := func(event *models.IndividualSymbolBookTicker) {
		logger.Infof("ticker: %s", event.String())
	}

	client := IndividualSymbolBookTickerStream(logger, symbols, handler)
	client.Run(ctx)

	time.Sleep(time.Second * 2) // 缩短测试时间
	client.Stop()
}

func TestKlinesStream(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要网络连接的测试")
	}

	ctx := context.TODO()
	logger := sugar.New(sugar.DefaultConfig())
	handler := func(event *klines.Data) {
		logger.Infof("klines: %+#v", event)
	}

	client := KlinesStream(logger, handler).Symbol("BTCUSDT").Interval("1m").TimeZone("UTC").Connect(ctx)
	client.Run(ctx)

	time.Sleep(time.Second * 10) // 缩短测试时间
	client.Stop()
}
