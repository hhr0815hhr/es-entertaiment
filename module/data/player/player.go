package player

import "github.com/gorilla/websocket"

type Player struct {
	Id       int64
	Username string
	Password string
	Nick     string
	Icon     string
	Coin     int64
	Level    int
	Sex      int
	Other    PlayerOther
}

type PlayerOther struct {
	Conn *websocket.Conn
}
