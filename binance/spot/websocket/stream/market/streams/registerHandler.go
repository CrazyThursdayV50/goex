package streams

import (
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/json"
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
	if handler != nil {
		goo.Go(func() {
			for data := range dataChan {
				handler(data)
			}
		})
	}

	s.dataChans[stream] = dataChan
}

func (s *Stream) HandleUnexpected(fn func([]byte)) {
	s.RegisterEventHandler("unexpected", fn)
}
