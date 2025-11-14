package utils

import (
	"fmt"
	"net/url"

	"github.com/CrazyThursdayV50/pkgo/json"
)

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

func MapString[T any](v T) (m map[string]string, err error) {
	var data []byte
	data, err = json.JSON().Marshal(v)
	if err != nil {
		return
	}

	err = json.JSON().Unmarshal(data, &m)
	return
}

func MapToQuery[T any](v T) (string, error) {
	m, err := MapAny(v)
	if err != nil {
		return "", err
	}

	var values url.Values
	for k, v := range m {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	return values.Encode(), nil
}

func JsonMarshalRaw[T any](v T) (json.RawMessage, error) {
	return json.JSON().Marshal(v)
}
