package handle

import (
	"context"
	"es-entertainment/common"
	"es-entertainment/lib/handle/proto"
	"fmt"
)

type DecodeStruct struct {
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data"`
}

type IDispatch interface {
	Handle(ctx context.Context, data []byte)
}

var (
	route map[int]IDispatch
)

func init() {
	//初始化dispatch map
	route = make(map[int]IDispatch)
	route[10001] = &proto.Login{}
	route[20000] = &proto.RoomList{}
	route[20001] = &proto.CreateRoom{}

}

func dispatch(ctx context.Context, data []byte) {
	//解码4字节的cmd 大端序
	if len(data) >= 4 {
		cmd := common.BytesToInt(data[:4])
		data = data[4:]
		err := common.CatchPanic(func() {
			route[cmd].Handle(ctx, data)
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
