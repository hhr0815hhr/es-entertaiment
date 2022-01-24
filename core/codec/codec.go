package codec

import (
	"es-entertainment/core/codec/json"
	"es-entertainment/core/codec/protobuf"
	"sync"
)

var (
	once      sync.Once
	codecType = CodeC_JSON
	c         ICodeC
)

func SetCodeCType(codeCType string) {
	codecType = codeCType
}

type ICodeC interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
}

func Instance() ICodeC {
	once.Do(func() {
		switch codecType {
		case CodeC_JSON:
			c = json.NewCodeC()
		case CodeC_PROTOBUF:
			c = protobuf.NewCodeC()
		}
	})
	return c
}
