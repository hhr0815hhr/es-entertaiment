package json

type JsonCodeC struct{}

func NewCodeC() *JsonCodeC {
	return &JsonCodeC{}
}

func (jc *JsonCodeC) Encode(data interface{}) ([]byte, error) {
	return nil, nil
}

func (jc *JsonCodeC) Decode(data []byte, v interface{}) error {
	return nil
}
