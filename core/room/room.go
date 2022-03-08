package room

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type Room struct {
	Id                 int32
	Type               string
	Name               string
	RoomPlayerLimitNum int32
	ChatChannel        chan string
	Players            []*RoomPlayer
	State              int32
	Timer              []interface{}
	Lock               *sync.RWMutex
}

type IRoom interface {
	Join(playerId int64) error
	Leave(playerId int64) error
	Run()
	Ready(playerId int64) error
}

var id uint32 = 100000

func NewRoom(roomName, roomType string, ctx context.Context) *Room {
	atomic.AddUint32(&id, 1)
	return &Room{
		Id:      int32(id),
		Type:    roomType,
		Name:    roomName,
		Players: make([]*RoomPlayer, 0),
		// RoomPlayerLimitNum: 2,
		ChatChannel: make(chan string, 500),
	}
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
	wg.Add(1)
	go r.chat(wg) // 开启聊天协程

	wg.Wait()
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
