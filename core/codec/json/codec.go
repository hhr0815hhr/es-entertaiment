package json

import "encoding/json"

type JsonCodeC struct{}

func NewCodeC() *JsonCodeC {
	return &JsonCodeC{}
}

func (jc *JsonCodeC) Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (jc *JsonCodeC) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
