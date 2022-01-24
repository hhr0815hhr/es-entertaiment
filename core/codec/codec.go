package codec

import (
	"errors"
	"es-entertainment/core/codec/json"
	"es-entertainment/core/codec/protobuf"
)

var CodeCType = CodeC_JSON

func SetCodeCType(codeCType string) {
	CodeCType = codeCType
}

type ICodeC interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
}

func NewCodeC(codecType string) (ICodeC, error) {
	switch codecType {
	case CodeC_JSON:
		return json.NewCodeC(), nil
	case CodeC_PROTOBUF:
		return protobuf.NewCodeC(), nil
	default:
		return nil, errors.New(ErrCodeCTypeNotFound)
	}
}
