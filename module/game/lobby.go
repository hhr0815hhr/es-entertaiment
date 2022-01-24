package game

import (
	"es-entertainment/core/room"
	"fmt"
	"sync"
)

type Lobby struct {
	Players          map[int]interface{} //sync.Map
	RoomManager      room.IRoomManager
	Lock             *sync.RWMutex
	LobbyChatChannel chan string
}

var LobbyInstance *Lobby

func InitLobby() {
	LobbyInstance = &Lobby{
		Players:          make(map[int]interface{}),
		RoomManager:      room.NewRoomManager(),
		Lock:             new(sync.RWMutex),
		LobbyChatChannel: make(chan string, 1000),
	}
	fmt.Println("初始化Lobby完成")
}

func (l *Lobby) GetRooms(roomType string) map[string]*room.Room {
	return l.RoomManager.GetRoomList(roomType)
}

func (l *Lobby) GetPlayers() map[int]interface{} {
	return l.Players
}

func (l *Lobby) EnterLobby(player interface{}) error {
	id := 2 //player.(*Player).GetId()
	l.Lock.Lock()
	defer l.Lock.Unlock()
	l.Players[id] = player
	l.LobbyChatChannel <- "欢迎玩家进入大厅"
	return nil
}

func (l *Lobby) LeaveLobby(player interface{}) error {
	id := 2 //player.(*Player).GetId()
	l.Lock.Lock()
	defer l.Lock.Unlock()
	delete(l.Players, id)
	l.LobbyChatChannel <- "玩家离开大厅"
	return nil
}
