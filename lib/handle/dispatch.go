package handle

import (
	"es-entertainment/common"
	"es-entertainment/lib/handle/proto"
)

type DecodeStruct struct {
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data"`
}

type IDispatch interface {
	Handle(data interface{})
}

var (
	route map[string]IDispatch
)

func init() {
	//初始化dispatch map
	route = make(map[string]IDispatch)
	route[`login`] = &proto.Login{}

}

func dispatch(cmd string, data interface{}) {
	common.CatchPanic(func() {
		route[cmd].Handle(data)
	})
}
