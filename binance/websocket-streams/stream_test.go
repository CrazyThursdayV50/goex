package websocketstreams

import (
	"context"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/websocket-streams/models"
	defaultlogger "github.com/CrazyThursdayV50/pkgo/log/default"
)

func TestStreamCreation(t *testing.T) {
	ctx := context.TODO()
	logger := defaultlogger.New(defaultlogger.DefaultConfig())
	logger.Init()
	
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
	logger := defaultlogger.New(defaultlogger.DefaultConfig())
	logger.Init()
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
	logger := defaultlogger.New(defaultlogger.DefaultConfig())
	logger.Init()
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
	logger := defaultlogger.New(defaultlogger.DefaultConfig())
	logger.Init()
	symbols := []string{"SOLBTC", "BNBBTC", "BNBUSDT"}
	handler := func(event *models.IndividualSymbolBookTicker) {
		logger.Infof("ticker: %s", event.String())
	}

	client := IndividualSymbolBookTickerStream(ctx, logger, symbols, handler)
	time.Sleep(time.Second * 2) // 缩短测试时间
	client.Stop()
}
