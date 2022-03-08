package room

import (
	"context"
	"errors"
	"es-entertainment/common"
	"fmt"
	"sync"
)

type Room struct {
	Id                 int
	Type               string
	Name               string
	RoomPlayerLimitNum int32
	ChatChannel        chan string
	Players            []RoomPlayer
	State              int
	Timer              []interface{}
	Lock               *sync.RWMutex
}

type IRoom interface {
	Join(player interface{}) error
	Leave(player interface{}) error
	Run()
	Ready(player interface{}) error
}

func NewRoom(roomName, roomType string, ctx context.Context) *Room {
	return &Room{
		// RoomPlayerLimitNum: 2,
		ChatChannel: make(chan string, 500),
	}
}

func (r *Room) Join(player interface{}) error {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	if ok, _ := common.Contain(player, r.Players); ok {
		return errors.New("player already in room")
	}
	if len(r.Players) >= int(r.RoomPlayerLimitNum) {
		return errors.New("room is full")
	}
	r.Players = append(r.Players, player.(RoomPlayer))
	return nil
}

func (r *Room) Leave(player interface{}) error {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	if ok, _ := common.Contain(player, r.Players); !ok {
		return errors.New("player not in room")
	}
	for i, p := range r.Players {
		if p == player {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			break
		}
	}
	return nil
}

func (r *Room) Ready(player interface{}) error {
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
