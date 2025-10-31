package utils

import "github.com/CrazyThursdayV50/pkgo/json"

func Ptr[T any](v T) *T { return &v }

func MapAny[T any](v T) (m map[string]any, err error) {
	var data []byte
	data, err = json.JSON().Marshal(v)
	if err != nil {
		return
	}

	err = json.JSON().Unmarshal(data, &m)
	return
}

func JsonMarshalRaw[T any](v T) (json.RawMessage, error) {
	return json.JSON().Marshal(v)
}
