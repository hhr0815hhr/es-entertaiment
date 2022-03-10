package room

import (
	"context"
	"errors"
	"es-entertainment/lib/chat"
	"es-entertainment/lib/send"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Room struct {
	Id                 int32
	Type               string
	Name               string
	RoomPlayerLimitNum int32
	CloseChannel       chan int
	ChatChannel        chan *chat.Chat
	// eventChannel       chan string
	EventTimer *time.Timer
	Tickers    map[string]*Ticker
	Players    []*RoomPlayer
	State      int32
	Ticker     []interface{}
	F          *RoomFsm
	Lock       *sync.RWMutex
}

type IRoom interface {
	Join(playerId int64) error
	Leave(playerId int64) error
	Run()
	Ready(player int64, state int) error
	Cast(playerId int64, cmd int32, msg []byte)
	Broadcast(cmd int32, msg []byte)
	// SetFsm(f *fsm.FSM)
}

var id uint32 = 100000

func NewRoom(roomName, roomType string, ctx context.Context) *Room {
	atomic.AddUint32(&id, 1)

	return &Room{
		Id:      int32(id),
		Type:    roomType,
		Name:    roomName,
		Players: make([]*RoomPlayer, 0),
		State:   0,
		// RoomPlayerLimitNum: 2,
		CloseChannel: make(chan int),
		ChatChannel:  make(chan *chat.Chat, 500),
		// eventChannel: make(chan string, 1),
		F: getRoomFsm(roomType),
	}
}

// func (r *Room) SetFsm(f *RoomFsm) {
// 	r.F = f
// }

func (r *Room) Cast(playerId int64, cmd int32, msg []byte) {
	send.SendToUid(playerId, msg, cmd)
}

func (r *Room) Broadcast(cmd int32, msg []byte) {
	// r.Lock.RLock()
	// defer r.Lock.RUnlock()
	uids := make([]int64, 0)
	for _, v := range r.Players {
		uids = append(uids, v.PlayerId)
	}
	send.SendToUids(uids, msg, cmd)
}

func (r *Room) Join(playerId int64) error {
	// r.Lock.Lock()
	// defer r.Lock.Unlock()
	for _, v := range r.Players {
		if v.PlayerId == playerId {
			return errors.New("player already in room")
		}
	}

	if len(r.Players) >= int(r.RoomPlayerLimitNum) {
		return errors.New("room is full")
	}
	roomPlayerInfo := initRoomPlayer(playerId)
	r.Players = append(r.Players, roomPlayerInfo)
	//退出大厅
	// game.LobbyInstance.LeaveLobby(playerId)
	return nil
}

func (r *Room) Leave(playerId int64) error {
	for i, p := range r.Players {
		if p.PlayerId == playerId {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			break
		}
	}
	if len(r.Players) == 0 {
		r.CloseChannel <- 1
	}
	return nil
}

func (r *Room) Ready(player int64, state int) error {
	for k, v := range r.Players {
		if v.PlayerId == player {
			r.Players[k].PlayerStatus = state
			break
		}
	}
	allReady := true
	for _, v := range r.Players {
		if v.PlayerStatus != 1 {
			allReady = false
			break
		}
	}
	if allReady {
		r.F.Event("start", r)
	}
	return nil
}

func (r *Room) Run() {
	// wg := new(sync.WaitGroup)
	// wg.Add(1)
	go r.chat() // 开启聊天协程
	// go r.running(wg)
	// wg.Wait()
	<-r.CloseChannel

}

// func (r *Room) running(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for x := range r.eventChannel {
// 		switch x {
// 		case "start":
// 			//启动game ticker
// 			//timer和event共同驱动fsm
// 		case "stop":
// 			//关闭game ticker
// 		default:
// 			log.Fatalf("unknown event: %s", x)
// 		}
// 	}

// }

func (r *Room) chat() {
	// defer wg.Done()
	for msg := range r.ChatChannel {
		fmt.Println(msg)
	}
	// for {
	// 	select {
	// 	case msg := <-r.ChatChannel:

	// 		// cast(r.Players,msg)

	// 	}
	// }
}
