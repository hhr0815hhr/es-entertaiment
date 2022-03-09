package room

import (
	"context"
	"errors"
	"es-entertainment/lib/chat"
	"es-entertainment/lib/send"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

type Room struct {
	Id                 int32
	Type               string
	Name               string
	RoomPlayerLimitNum int32
	ChatChannel        chan *chat.Chat
	eventChannel       chan string
	Players            []*RoomPlayer
	State              int32
	Ticker             []interface{}
	F                  *Fsm
	Lock               *sync.RWMutex
}

type IRoom interface {
	Join(playerId int64) error
	Leave(playerId int64) error
	Run()
	Ready(playerId int64) error
	Cast(playerId int64, cmd int32, msg []byte)
	Broadcast(cmd int32, msg []byte)
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
		ChatChannel:  make(chan *chat.Chat, 500),
		eventChannel: make(chan string, 1),
		F:            initFsm(roomType, getFsm(roomType)), //注入fsm,
	}
}


func (r *Room) Cast(playerId int64, cmd int32, msg []byte) {
	send.SendToUid(playerId, msg, cmd)
}

func (r *Room) Broadcast(cmd int32, msg []byte) {
	r.Lock.RLock()
	defer r.Lock.RUnlock()
	uids := make([]int64, 0)
	for _, v := range r.Players {
		uids = append(uids, v.PlayerId)
	}
	send.SendToUids(uids, msg, cmd)
}

func (r *Room) Join(playerId int64) error {
	r.Lock.Lock()
	defer r.Lock.Unlock()
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
	r.Lock.Lock()
	defer r.Lock.Unlock()

	// if ok, _ := common.Contain(player, r.Players); !ok {
	// 	return errors.New("player not in room")
	// }
	for i, p := range r.Players {
		if p.PlayerId == playerId {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			break
		}
	}
	return nil
}

func (r *Room) Ready(player int64) error {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	return nil
}

func (r *Room) Run() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go r.chat(wg) // 开启聊天协程
	go r.running(wg)
	wg.Wait()
}

func (r *Room) running(wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range r.eventChannel {
		switch x {
		case "start":
			//启动game ticker
			//timer和event共同驱动fsm
		case "stop":
			//关闭game ticker
		default:
			log.Fatalf("unknown event: %s", x)
		}
	}

}

func (r *Room) chat(wg *sync.WaitGroup) {
	defer wg.Done()
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
