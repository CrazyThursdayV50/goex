package streams

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/worker"
	"github.com/tidwall/gjson"
)

type Stream struct {
	id       int64
	logger   log.Logger
	wsclient *client.Client

	workers   map[string]*worker.Worker[[]byte]
	dataChans map[string]chan []byte
}

func (s *Stream) Clone() *Stream {
	var stream Stream
	stream.logger = s.logger
	stream.workers = make(map[string]*worker.Worker[[]byte])
	stream.dataChans = make(map[string]chan []byte)
	stream.HandleUnexpected(func(b []byte) {
		stream.logger.Warnf("unexpected data: %s", b)
	})

	return &stream
}

func (s *Stream) SetReqID(req interface{ SetID(int64) }) int64 {
	id := atomic.AddInt64(&s.id, 1)
	req.SetID(id)
	return id
}

func (s *Stream) sendEventData(ctx context.Context, event string, data []byte) {
	ch, ok := s.dataChans[event]
	if !ok {
		s.logger.Warnf("event chan not found: %s", event)
		return
	}

	select {
	case <-ctx.Done():
	case ch <- data:
	}
}

func (s *Stream) handler(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
	// 尝试解析为订阅的返回值
	result := gjson.GetBytes(b, "id")
	if result.Exists() {
		var resp models.Response
		err := json.JSON().Unmarshal(b, &resp)
		if err != nil {
			f(fmt.Errorf("unmarshal response failed: %v", err))
			return client.TextMessage, nil
		}

		// 订阅消息暂时就打印一下
		l.Debugf("Subscribe reponse: %+v", resp)
		return client.TextMessage, nil
	}

	results := gjson.GetManyBytes(b, "stream", "data")
	// 是 combined stream 消息
	if results[0].Exists() {
		s.sendEventData(ctx, results[0].String(), []byte(results[1].String()))
		return client.TextMessage, nil
	}

	// 不是 combined stream 消息
	// 尝试解析其中的 event
	result = gjson.GetBytes(b, "e")
	if result.Exists() {
		s.sendEventData(ctx, result.String(), b)
		return client.TextMessage, nil
	}

	result = gjson.GetBytes(b, "0.e")
	if result.Exists() {
		s.sendEventData(ctx, result.String(), b)
		return client.TextMessage, nil
	}

	l.Warnf("unexpected data: %s", b)
	return client.TextMessage, nil
}

func New(logger log.Logger) *Stream {
	var stream Stream
	stream.dataChans = make(map[string]chan []byte)
	stream.workers = make(map[string]*worker.Worker[[]byte])
	stream.logger = logger
	stream.HandleUnexpected(func(b []byte) {
		logger.Warnf("unexpected data: %s", b)
	})

	return &stream
}

func (s *Stream) Run(ctx context.Context) error {
	for _, w := range s.workers {
		w.Run(ctx)
	}
	return s.wsclient.Run(ctx)
}

func request[reqParams any](stream *Stream, req *models.Request[reqParams]) error {
	stream.SetReqID(req)

	data, err := json.JSON().Marshal(req)
	if err != nil {
		return err
	}

	return stream.wsclient.Send(data)
}

func (s *Stream) OnConnect(f func() (int, []byte)) *Stream {
	s.wsclient.OnConnect(f)
	return s
}
