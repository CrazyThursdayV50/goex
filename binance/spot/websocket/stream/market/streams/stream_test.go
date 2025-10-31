package streams

import (
	"context"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models/bookticker"
	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/stream/market/models/depth"
	"github.com/CrazyThursdayV50/pkgo/builtin/collector"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type testLogger struct {
	t *testing.T
}

func (l *testLogger) Debug(args ...any) {
	l.t.Log(args...)
}

func (l *testLogger) Info(args ...any) {
	l.t.Log(args...)
}

func (l *testLogger) Warn(args ...any) {
	l.t.Log(args...)
}

func (l *testLogger) Error(args ...any) {
	l.t.Error(args...)
}

func (l *testLogger) Debugf(fmt string, args ...any) {
	l.t.Logf(fmt, args...)
}

func (l *testLogger) Infof(fmt string, args ...any) {
	l.t.Logf(fmt, args...)
}

func (l *testLogger) Warnf(fmt string, args ...any) {
	l.t.Logf(fmt, args...)
}

func (l *testLogger) Errorf(fmt string, args ...any) {
	l.t.Errorf(fmt, args...)
}

type StreamEnv struct {
	symbol  string
	symbols []string
	logger  log.Logger

	bookDepthLevel int

	stream *Stream
}

func Setup(t *testing.T) *StreamEnv {
	var env = StreamEnv{
		symbol:         "BTCUSDT",
		symbols:        []string{"SOLBTC", "BNBBTC", "BNBUSDT"},
		logger:         &testLogger{t},
		bookDepthLevel: 10,
	}

	env.stream = New(env.logger)
	return &env
}

func Test(t *testing.T) {
	env := Setup(t)
	// t.Run("IndividualSymbol BookTicker", env.TestStreamBookTicker)
	// t.Run("Combined Individual BookTicker", env.TestCombinedBookTicker)
	t.Run("PartialBookDepth", env.TestStreamPartialBookDepth)
}

func (env *StreamEnv) handlerBookTicker() WsIndividualSymbolBookTickerHandler {
	return func(event *bookticker.IndividualSymbolBookTicker, err error) {
		if err != nil {
			env.logger.Errorf("handle bookticker failed: %v", err)
			return
		}

		env.logger.Infof("bookticker: %s", event.String())
	}
}

func (env *StreamEnv) handlerPartialBookDepth() WsPartialDepthHandler {
	return func(event *depth.PartialDepthData, err error) {
		if err != nil {
			env.logger.Errorf("handle partial depth failed: %v", err)
			return
		}

		env.logger.Infof("partial depth: %s", event.String())
	}
}

func (env *StreamEnv) TestStreamBookTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	streamName := bookticker.StreamName(env.symbol)
	stream := env.stream.IndividualBookTicker(streamName, env.handlerBookTicker())
	stream.Run(ctx)
	time.Sleep(time.Second * 2) // 缩短测试时间
}

func (env *StreamEnv) TestCombinedBookTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	streamNames := collector.Slice(env.symbols, func(_ int, symbol string) (bool, string) {
		return true, bookticker.StreamName(symbol)
	})

	stream := env.stream.Combined(streamNames)
	for _, streamName := range streamNames {
		stream.HandleBookTicker(streamName, env.handlerBookTicker())
	}

	stream.Run(ctx)
	time.Sleep(time.Second * 2)
}

func (env *StreamEnv) TestStreamPartialBookDepth(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	streamName := depth.StreamName(env.symbol, env.bookDepthLevel)
	stream := env.stream.PartialBookDepth5Stream(streamName, env.handlerPartialBookDepth())
	stream.Run(ctx)
	time.Sleep(time.Second * 2) // 缩短测试时间
}

// func TestStreamCreation(t *testing.T) {

// 	// 测试创建流客户端但不连接
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Errorf("创建流客户端时发生panic: %v", r)
// 		}
// 	}()

// 	handler := func(event *models.PartialDepthData) {
// 		logger.Infof("depth: %s", event.String())
// 	}

// 	// 创建客户端但立即停止，只测试创建过程
// 	client := PartialBookDepth5Stream(ctx, logger, symbol, handler)
// 	if client == nil {
// 		t.Errorf("创建流客户端失败")
// 		return
// 	}

// 	// 立即停止，不等待连接
// 	client.Stop()
// 	t.Log("流客户端创建测试通过")
// }

// func TestStream(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("跳过需要网络连接的测试")
// 	}

// 	ctx := context.TODO()
// 	logger := sugar.New(sugar.DefaultConfig())
// 	symbol := "BTCUSDT"
// 	handler := func(event *models.PartialDepthData) {
// 		logger.Infof("depth: %s", event.String())
// 	}

// 	client := PartialBookDepth5Stream(ctx, logger, symbol, handler)
// 	time.Sleep(time.Second * 2) // 缩短测试时间
// 	client.Stop()
// }

// func TestCombinedStream(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("跳过需要网络连接的测试")
// 	}

// 	ctx := context.TODO()
// 	logger := sugar.New(sugar.DefaultConfig())
// 	symbols := []string{"SOLBTC", "BNBBTC", "BNBUSDT"}
// 	handler := func(event *models.PartialDepthCombinedData) {
// 		logger.Infof("depth: %s", event.String())
// 	}

// 	client := PartialBookDepth5CombinedStream(ctx, logger, symbols, handler)
// 	time.Sleep(time.Second * 2) // 缩短测试时间
// 	client.Stop()
// }

// func TestKlinesStream(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("跳过需要网络连接的测试")
// 	}

// 	ctx := context.TODO()
// 	logger := sugar.New(sugar.DefaultConfig())
// 	handler := func(event *klines.Data) {
// 		logger.Infof("klines: %+#v", event)
// 	}

// 	client := KlinesStream(logger, handler).Symbol("BTCUSDT").Interval("1m").TimeZone("UTC").Connect(ctx)
// 	client.Run(ctx)

// 	time.Sleep(time.Second * 10) // 缩短测试时间
// 	client.Stop()
// }
