package websocketstreams

import (
	"context"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/models"
	defaultlogger "github.com/CrazyThursdayV50/pkgo/log/default"
)

func TestStream(t *testing.T) {
	ctx := context.TODO()
	logger := defaultlogger.New(defaultlogger.DefaultConfig())
	logger.Init()
	symbol := "BTCUSDT"
	handler := func(event *models.PartialDepthData) {
		logger.Infof("depth: %s", event.String())
	}

	client := PartialBookDepth5Stream(ctx, logger, symbol, handler)
	time.Sleep(time.Second * 10)
	client.Stop()
}

func TestCombinedStream(t *testing.T) {
	ctx := context.TODO()
	logger := defaultlogger.New(defaultlogger.DefaultConfig())
	logger.Init()
	symbols := []string{"SOLBTC", "BNBBTC", "BNBUSDT"}
	handler := func(event *models.PartialDepthCombinedData) {
		logger.Infof("depth: %s", event.String())
	}

	client := PartialBookDepth5CombinedStream(ctx, logger, symbols, handler)
	time.Sleep(time.Second * 10)
	client.Stop()
}
