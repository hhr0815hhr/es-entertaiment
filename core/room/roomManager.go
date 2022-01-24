package room

import (
	"context"
	"sync"
)

type RoomManager struct {
	rooms map[string]map[string]*Room
	Lock  *sync.RWMutex
}

type IRoomManager interface {
	CreateRoom(roomName, roomType string) *Room
	DestroyRoom(roomName string)
	GetRoomList(roomType string) map[string]*Room
}

var RoomManagerInstance *RoomManager
var once sync.Once

func NewRoomManager() IRoomManager {
	once.Do(func() {
		RoomManagerInstance = &RoomManager{
			rooms: make(map[string]map[string]*Room),
			Lock:  new(sync.RWMutex),
		}
	})
	return RoomManagerInstance
}

func (rm *RoomManager) CreateRoom(roomName, roomType string) *Room {
	rm.Lock.Lock()
	defer rm.Lock.Unlock()
	ctx, _ := context.WithCancel(context.Background())
	room := NewRoom(roomName, roomType, ctx)
	rm.rooms[roomType][roomName] = room
	go room.Run()
	return room
}

func (rm *RoomManager) DestroyRoom(roomName string) {
	rm.Lock.Lock()
	defer rm.Lock.Unlock()
	delete(rm.rooms, roomName)
}

func (rm *RoomManager) GetRoomList(roomType string) map[string]*Room {
	rm.Lock.RLock()
	defer rm.Lock.RUnlock()
	return rm.rooms[roomType]
}
