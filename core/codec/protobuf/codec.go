package protobuf

import "es-entertainment/common"

type ProtobufCodeC struct{}

func NewCodeC() *ProtobufCodeC {
	return &ProtobufCodeC{}
}

func (jc *ProtobufCodeC) Encode(data interface{}) ([]byte, error) {
	return nil, nil
}

func (jc *ProtobufCodeC) Decode(data []byte, v interface{}) error {
	//解码4字节的cmd 大端序
	cmd := common.BytesToInt(data[:4])
	data = data[4:]

	// name := "logon.SearchRequest"
	// pt := proto.MessageType(name)
	// a := reflect.New(pt.Elem()).Interface().(proto.Message)
	// proto.Unmarshal(dst, a)
	// fmt.Printf("方式二 ：unmarshaled message: %v\n", a)

	return nil
}
