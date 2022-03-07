package handle

import (
	"context"
	"encoding/json"
	"es-entertainment/common"
	"es-entertainment/core/codec"
	"es-entertainment/lib/handle/proto"
	"fmt"
	"reflect"
)

type DecodeStruct struct {
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data"`
}

type IDispatch interface {
	Handle(ctx context.Context, data interface{})
}

var (
	route map[int]IDispatch
)

func init() {
	//初始化dispatch map
	route = make(map[int]IDispatch)
	route[10001] = &proto.Login{}
	route[20001] = &proto.CreateRoom{}
}

func dispatch(ctx context.Context, data []byte) {

	var decode DecodeStruct
	err := codec.Instance().Decode(data, &decode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解码数据：", decode)
	cmd, dData := decode.Cmd, decode.Data
	err = common.CatchPanic(func() {
		json.Unmarshal(dData.([]byte), reflect.TypeOf(route[cmd]))
		route[cmd].Handle(ctx, dData)
	})
	if err != nil {
		fmt.Println(err)
	}
}
