package proto

import (
	"es-entertainment/protos"
	"testing"
)

type Login protos.C2S_Login

type L struct {
	C *protos.C2S_Login
}

func TestProto(t *testing.T) {

	// a := "protos.C2S_Login"
	l := &Login{}
	interface{}(l).(*protos.C2S_Login).Marshal()

	l2 := &L{}
	l2.C.Marshal()
}
