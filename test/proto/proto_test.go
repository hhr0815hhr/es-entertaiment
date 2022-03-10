package proto

import (
	"es-entertainment/core/codec"
	"es-entertainment/protos"
	"fmt"
	"testing"
)

type Login protos.C2S_Login

type L struct {
	C *protos.C2S_Login
}

func TestProto(t *testing.T) {

	// a := "protos.C2S_Login"
	a := &protos.C2S_CreateRoom{
		RoomName: "test",
		RoomType: "cow",
	}
	fmt.Println(a)
	ab, _ := codec.Instance().Encode(a)
	fmt.Println(ab)
	c := &protos.C2S_CreateRoom{}
	codec.Instance().Decode(ab, c)
	fmt.Println(c)
	// t.Errorf("%v", c == a)
}
