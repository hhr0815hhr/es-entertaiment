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
	route map[string]IDispatch
)

func init() {
	//初始化dispatch map
	route = make(map[string]IDispatch)
	route[`login`] = &proto.Login{}
	route[`createRoom`] = &proto.CreateRoom{}
}

func dispatch(ctx context.Context, data interface{}) {
	//解码
	var decode DecodeStruct
	err := codec.Instance().Decode(data.([]byte), &decode)
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
