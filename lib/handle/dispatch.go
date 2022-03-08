package handle

import (
	"context"
	"es-entertainment/common"
	"es-entertainment/lib/handle/proto"
	"es-entertainment/protos"
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
	route[int(protos.CmdType_CMD_C2S_Login)] = &proto.Login{}
	route[int(protos.CmdType_CMD_C2S_RoomList)] = &proto.RoomList{}
	route[int(protos.CmdType_CMD_C2S_CreateRoom)] = &proto.CreateRoom{}
	route[int(protos.CmdType_CMD_C2S_EnterRoom)] = &proto.EnterRoom{}
	route[int(protos.CmdType_CMD_C2S_LeaveRoom)] = &proto.LeaveRoom{}
	route[int(protos.CmdType_CMD_C2S_Ready)] = &proto.Ready{}

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
