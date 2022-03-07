package server

import (
	"errors"
	"es-entertainment/core/server/tcp"
	"es-entertainment/core/server/udp"
	"es-entertainment/core/server/websocket"
)

type IServer interface {
	BeforeRun()
	Run(host string, port int) error
	Stop()
	SetHandle(f interface{}, fHttp interface{})
}

func NewServer(id int64, serverType string) (IServer, error) {
	switch serverType {
	case Server_Type_TCP:
		return tcp.NewServer(id), nil
	case Server_Type_UDP:
		return udp.NewServer(id), nil
	case Server_Type_WS:
		return websocket.NewServer(id), nil
	default:
		return nil, errors.New("server type not support")
	}
}
