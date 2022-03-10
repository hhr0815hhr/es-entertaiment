package bytes

import (
	"es-entertainment/common"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	// var buf = b.NewBuffer(nil)
	// // var buf2 = b.NewBuffer(nil)
	// cmd := int32(10001)
	// buf.WriteByte(byte(cmd >> 24))
	// buf.WriteByte(byte(cmd >> 16))
	// buf.WriteByte(byte(cmd >> 8))
	// buf.WriteByte(byte(cmd))

	buf := common.IntToBytes(10001)
	fmt.Printf("%v", buf)
}
