package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type ProtobufCodeC struct{}

func NewCodeC() *ProtobufCodeC {
	return &ProtobufCodeC{}
}

func (jc *ProtobufCodeC) Encode(data interface{}) ([]byte, error) {
	return proto.Marshal(data.(proto.Message))
}

func (jc *ProtobufCodeC) Decode(data []byte, v interface{}) error {

	//proto
	// reflect.TypeOf(v)
	// proto.MessageName(v)
	return proto.Unmarshal(data, v.(proto.Message))
}
