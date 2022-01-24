package protobuf

type ProtobufCodeC struct{}

func NewCodeC() *ProtobufCodeC {
	return &ProtobufCodeC{}
}

func (jc *ProtobufCodeC) Encode(data interface{}) ([]byte, error) {
	return nil, nil
}

func (jc *ProtobufCodeC) Decode(data []byte, v interface{}) error {
	return nil
}
