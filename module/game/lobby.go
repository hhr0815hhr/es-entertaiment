package game

import (
	"es-entertainment/common"
	"es-entertainment/core/codec"
	"es-entertainment/core/log"
	"es-entertainment/core/room"
	"es-entertainment/lib/send"
	"es-entertainment/module/data/player"
	"es-entertainment/protos"
	"sync"
	"time"
)

type Lobby struct {
	Players          map[int64]*LobbyPlayer //sync.Map
	RoomManager      room.IRoomManager
	Lock             *sync.RWMutex
	LobbyChatChannel chan *LobbyChat
}

type LobbyPlayer struct {
	Id   int64
	Nick string
	Icon string
}

type LobbyChat struct {
	From int64
	Msg  string
	Time int64
}

var LobbyInstance *Lobby

func InitLobby() {
	LobbyInstance = &Lobby{
		Players:          make(map[int64]*LobbyPlayer),
		RoomManager:      room.NewRoomManager(),
		Lock:             new(sync.RWMutex),
		LobbyChatChannel: make(chan *LobbyChat, 1000),
	}
	go lobbyChat(LobbyInstance)
	log.Info("init lobby success...")
}

func (l *Lobby) GetRooms(roomType string) map[int32]*room.Room {
	return l.RoomManager.GetRoomList(roomType)
}

func (l *Lobby) GetPlayers() map[int64]*LobbyPlayer {
	return l.Players
}

func (l *Lobby) GetPlayerUids() []int64 {
	uids := make([]int64, len(l.Players))
	for _, v := range l.Players {
		uids = append(uids, v.Id)
	}
	return uids
}

func (l *Lobby) EnterLobby(p interface{}) error {
	info := p.(*player.Player) //player.(*Player).GetId()
	lobbyPlayer := &LobbyPlayer{
		Id:   info.Id,
		Nick: info.Nick,
		Icon: info.Icon,
	}
	l.Lock.Lock()
	defer l.Lock.Unlock()
	l.Players[info.Id] = lobbyPlayer
	msg := &LobbyChat{
		From: 0,
		Msg:  "玩家:" + info.Nick + "进入大厅",
		Time: time.Now().Unix(),
	}
	l.LobbyChatChannel <- msg
	return nil
}

func (l *Lobby) BroadCast(msg *LobbyChat) {
	pb := &protos.S2C_Chat{
		From:    msg.From,
		Msg:     msg.Msg,
		Time:    msg.Time,
		Channel: player.Position_Lobby,
	}
	b, _ := codec.Instance().Encode(pb)
	send.SendToUids(l.GetPlayerUids(), b, int32(protos.CmdType_CMD_S2C_Chat))
}

func (l *Lobby) LeaveLobby(p interface{}) error {
	info := p.(*player.Player)
	l.Lock.Lock()
	defer l.Lock.Unlock()
	delete(l.Players, info.Id)
	msg := &LobbyChat{
		From: 0,
		Msg:  "玩家:" + info.Nick + "离开大厅",
		Time: time.Now().Unix(),
	}
	l.LobbyChatChannel <- msg
	return nil
}

func lobbyChat(lobby *Lobby) {
	common.RunNoPanic(func() {
		for msg := range lobby.LobbyChatChannel {
			lobby.BroadCast(msg)
			// fmt.Printf("大厅推送消息：%s", msg)
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
