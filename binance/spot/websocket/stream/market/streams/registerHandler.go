package streams

import (
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/worker"
)

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

	worker, _ := worker.New(stream, handler)
	worker.WithTrigger(dataChan)
	worker.WithGraceful(true)

	s.workers[stream] = worker
	s.dataChans[stream] = dataChan
}

func (s *Stream) HandleUnexpected(fn func([]byte)) {
	s.RegisterEventHandler("unexpected", fn)
}
