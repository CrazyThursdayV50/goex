package streams

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/CrazyThursdayV50/goex/binance/variables"
	"github.com/CrazyThursdayV50/goex/binance/variables/derivatives"
	"github.com/CrazyThursdayV50/goex/infra/websocket/client"
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

type Stream struct {
	id        int64
	logger    log.Logger
	wsclient  *client.Client
	listenkey string

	l         sync.RWMutex
	dataChans map[string]chan []byte

	done chan struct{}
}

func (s *Stream) SetListenKey(listenkey string) {
	s.listenkey = listenkey
}

func (s *Stream) Clone() *Stream {
	var stream Stream
	stream.listenkey = s.listenkey
	stream.logger = s.logger
	stream.dataChans = make(map[string]chan []byte)
	stream.done = make(chan struct{})
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

func (s *Stream) sendEventData(event string, data []byte) {
	s.l.RLock()
	ch, ok := s.dataChans[event]
	s.l.RUnlock()

	if !ok {
		s.logger.Warnf("event chan not found: %s", event)
		return
	}

	select {
	case <-s.done:
		close(ch)

		s.l.Lock()
		delete(s.dataChans, event)
		s.l.Unlock()

	case ch <- data:
	}
}

func (s *Stream) handler(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
	// 不是 combined stream 消息
	// 尝试解析其中的 event
	result := gjson.GetBytes(b, "e")
	if result.Exists() {
		s.sendEventData(result.String(), b)
		return client.TextMessage, nil
	}

	l.Warnf("unexpected data: %s", b)
	return client.TextMessage, nil
}

func New(logger log.Logger, listenKey string) *Stream {
	var stream Stream
	stream.listenkey = listenKey
	stream.dataChans = make(map[string]chan []byte)
	stream.done = make(chan struct{})
	stream.logger = logger
	stream.HandleUnexpected(func(b []byte) {
		logger.Warnf("unexpected data: %s", b)
	})

	stream.wsclient = client.NewClient(
		stream.logger,
		fmt.Sprintf("%s/%s", derivatives.WsStream().Endpoint(), stream.listenkey),
		variables.GetProxy(),
		stream.handler,
		func(done <-chan struct{}, conn *websocket.Conn) {
			ticker := time.NewTicker(time.Minute * 15)
			for {
				select {
				case <-done:
					return

				case t := <-ticker.C:
					conn.WriteControl(
						client.PingMessage,
						fmt.Appendf(nil, "%d", t.UnixMilli()),
						time.Now().Add(time.Second*30))

					conn.WriteControl(
						client.PongMessage,
						fmt.Appendf(nil, "%d", t.UnixMilli()),
						time.Now().Add(time.Second*30))
				}
			}
		},
	)

	return &stream
}

func (s *Stream) Run(ctx context.Context) error {
	err := s.wsclient.Run(ctx)
	if err != nil {
		return err
	}

	goo.Go(func() {
		<-s.done
		s.wsclient.Stop()
	})

	return nil
}

func (s *Stream) Stop() {
	close(s.done)
}

func (s *Stream) OnConnect(f func() (int, []byte)) *Stream {
	s.wsclient.OnConnect(f)
	return s
}
