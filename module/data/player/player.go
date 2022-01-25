package player

import "github.com/gorilla/websocket"

type Player struct {
	Id       int64
	Username string
	Other    PlayerOther
}

type PlayerOther struct {
	Conn *websocket.Conn
}
