package pack

import (
	"es-entertainment/core/pack/common"
	"es-entertainment/core/pack/eof"
	"sync"
)

type IPack interface {
	Pack(data interface{}) ([]byte, error)
	Unpack(data []byte) (interface{}, error)
}

const (
	Pack_Type_EOF = "eof"
	Pack_Type_CM  = "common"
)

var (
	once     sync.Once
	p        IPack
	packType = Pack_Type_EOF
)

func SetPackType(pt string) {
	packType = pt
}

func Instance() IPack {
	once.Do(func() {
		switch packType {
		case Pack_Type_EOF:
			p = eof.NewPack()
		case Pack_Type_CM:
			p = common.NewPack()
		}
	})
	return p
}
