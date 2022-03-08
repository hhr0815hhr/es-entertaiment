package game

import (
	"es-entertainment/common"
	"es-entertainment/core/log"
	"es-entertainment/core/room"
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
	go lobbyChat(LobbyInstance)
	log.Info("init lobby success...")
}

func (l *Lobby) GetRooms(roomType string) map[int32]*room.Room {
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

func lobbyChat(lobby *Lobby) {
	common.RunNoPanic(func() {
		for msg := range lobby.LobbyChatChannel {
			log.Infof("大厅推送消息：%s", msg)
		}
		// for {
		// 	select {
		// 	case msg := <-lobby.LobbyChatChannel:
		// 		log.Infof("大厅推送消息：%s", msg)
		// 		//对所有大厅玩家推送消息
		// 		// for k, v := range lobby.Players {
		// 		// 	fmt.Println(k, v)
		// 		// 	//v.conn.Write([]byte(msg))
		// 		// }
		// 	}
		// }
	})
}
