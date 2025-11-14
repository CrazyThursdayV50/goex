package market

import (
	"context"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/klines"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/klines/continuous"
	markprice "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/markPrice"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models/subscribe"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/streams"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
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
	symbol   string
	interval string

	symbols []string
	logger  log.Logger

	bookDepthLevel int
	contract       continuous.ContractType

	stream *streams.Stream
}

func Setup(t *testing.T) *StreamEnv {
	var env = StreamEnv{
		symbol:   "BTCUSDT",
		interval: "1m",

		symbols: []string{"SOLBTC", "BNBBTC", "BNBUSDT"},
		logger:  &testLogger{t},

		contract:       continuous.Perpetual,
		bookDepthLevel: 10,
	}

	env.stream = streams.New(env.logger)
	return &env
}

func handleStream[Data any](env *StreamEnv, name string) func(Data, error) {
	return func(d Data, err error) {
		if err != nil {
			env.logger.Errorf("%s failed: %v", name, err)
			return
		}
		env.logger.Debugf("%s: %+v", name, d)
	}
}

func TestStream(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	env := Setup(t)
	stream := env.stream.Stream()
	connected := make(chan struct{})

	stream.OnConnect(func() (int, []byte) {
		close(connected)
		return client.TextMessage, nil
	})

	stream.HandleKlineStreamEvent(handleStream[*klines.Result](env, klines.Event))
	stream.HandleContinuousKlineStreamEvent(handleStream[*klines.Result](env, klines.Event))
	stream.HandleMarkPriceStreamEvent(handleStream[*markprice.Result](env, markprice.Event))
	stream.Run(ctx)
	<-connected

	streams := subscribe.RequestParams{
		klines.StreamName(env.symbol, env.interval),
		continuous.StreamName(env.symbol, env.contract.String(), env.interval),
		markprice.StreamName(env.symbol),
	}

	t.Run("Subscribe", func(t *testing.T) {
		err := stream.Subscribe(t.Context(), streams)

		if err != nil {
			t.Fatalf("subscribe failed: %v", err)
		}

		time.Sleep(time.Second * 3)
	})

	t.Run("Unsubscribe", func(t *testing.T) {
		err := stream.Unsubscribe(t.Context(), streams)
		if err != nil {
			t.Fatalf("unsubscribe failed: %v", err)
		}

		time.Sleep(time.Second * 3)
	})

}

func TestCombined(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	env := Setup(t)
	stream := env.stream.Combined()
	connected := make(chan struct{})

	stream.OnConnect(func() (int, []byte) {
		close(connected)
		return client.TextMessage, nil
	})

	stream.HandleKlineCombinedData(
		klines.StreamName(env.symbol, env.interval),
		handleStream[*klines.Result](env, klines.StreamName(env.symbol, env.interval)),
	)

	stream.HandleContinuousKlineCombinedData(
		continuous.StreamName(env.symbol, env.contract.String(), env.interval),
		handleStream[*klines.Result](env, continuous.StreamName(env.symbol, env.contract.String(), env.interval)),
	)

	stream.HandleMarkPriceCombinedData(markprice.StreamName(env.symbol), handleStream[*markprice.Result](env, markprice.StreamName(env.symbol)))

	stream.Run(ctx)
	<-connected

	streams := subscribe.RequestParams{
		klines.StreamName(env.symbol, env.interval),
		continuous.StreamName(env.symbol, env.contract.String(), env.interval),
		markprice.StreamName(env.symbol),
	}

	t.Run("Subscribe", func(t *testing.T) {
		err := stream.Subscribe(t.Context(), streams)

		if err != nil {
			t.Fatalf("subscribe failed: %v", err)
		}

		time.Sleep(time.Second * 3)
	})

	t.Run("Unsubscribe", func(t *testing.T) {
		err := stream.Unsubscribe(t.Context(), streams)

		if err != nil {
			t.Fatalf("subscribe failed: %v", err)
		}

		time.Sleep(time.Second * 3)
	})
}
