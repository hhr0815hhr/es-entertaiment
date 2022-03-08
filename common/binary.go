package common

import (
	"bytes"
	"encoding/binary"
)

//整形转换成字节
func IntToBytes(n int) []byte {
	// var buf = bytes.NewBuffer(nil)
	// buf.WriteByte(byte(cmd << 24))
	// buf.WriteByte(byte(cmd << 16))
	// buf.WriteByte(byte(cmd << 8))
	// buf.WriteByte(byte(cmd))
	// buf.Write(data)
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
