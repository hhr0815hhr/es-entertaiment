package client

import "github.com/gorilla/websocket"

type CliConn *websocket.Conn

type Client struct {
	Id   int64
	Conn CliConn
}
