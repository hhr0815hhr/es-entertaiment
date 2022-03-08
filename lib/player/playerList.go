package player

import (
	"sync"

	"github.com/gorilla/websocket"
)

type PlayerList struct {
	*sync.Map
}

type Player struct {
	Id    int
	Nick  string
	Icon  string
	Coin  int
	Level int
	Sex   int
	Other PlayerOther
}

type PlayerOther struct {
	Conn *websocket.Conn
}
