package streams

import (
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models/account"
	listenkey "github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models/listenKey"
	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/user/models/order"
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/json"
)

func wrapHandler[T any](first, second func(T, error)) func(T, error) {
	return func(t T, err error) {
		if first != nil {
			first(t, err)
		}

		if second != nil {
			second(t, err)
		}
	}
}

func CreateBytesHandler[T any](fn func(T, error)) func([]byte) {
	return func(b []byte) {
		var result T
		err := json.JSON().Unmarshal(b, &result)
		fn(result, err)
	}
}

// 注册不同 event 的handler
func (s *Stream) RegisterEventHandler(stream string, handler func([]byte)) {
	var dataChan = make(chan []byte, 100)
	if handler != nil {
		goo.Go(func() {
			for {
				select {
				case <-s.done:
					return

				case data := <-dataChan:
					handler(data)
				}
			}
		})
	}

	s.l.Lock()
	defer s.l.Unlock()
	s.dataChans[stream] = dataChan
}

func (s *Stream) HandleUnexpected(fn func([]byte)) {
	s.RegisterEventHandler("unexpected", fn)
}

func (s *Stream) HandleListenKeyExpired(handler func(*listenkey.Event, error)) {
	handler = wrapHandler(handler, func(event *listenkey.Event, err error) {
		if event != nil {
			s.Stop()
		}
	})

	s.RegisterEventHandler(
		listenkey.EventName,
		CreateBytesHandler(handler),
	)
}

func (s *Stream) HandleAccountUpdate(handler func(*account.Event, error)) {
	s.RegisterEventHandler(
		account.EventName,
		CreateBytesHandler(handler),
	)
}

func (s *Stream) HandleOrderUpdate(handler func(*order.Event, error)) {
	s.RegisterEventHandler(
		order.EventName,
		CreateBytesHandler(handler),
	)
}
