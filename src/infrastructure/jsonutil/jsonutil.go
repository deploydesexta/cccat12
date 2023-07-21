package jsonutil

import (
	"github.com/goccy/go-json"
	"io"
)

func ToJsonBytes(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func ToJson(i interface{}) (string, error) {
	b, err := ToJsonBytes(i)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func FromJson(data []byte, i interface{}) error {
	return json.Unmarshal(data, i)
}

func FromReader(r io.Reader, i interface{}) error {
	return json.NewDecoder(r).Decode(i)
}
