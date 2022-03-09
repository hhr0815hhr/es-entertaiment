package room

import (
	"context"
	"sync"
)

type RoomManager struct {
	rooms map[string]map[int32]*Room
	ctx   context.Context
	Lock  *sync.RWMutex
}

type IRoomManager interface {
	CreateRoom(roomName, roomType string) *Room
	DestroyRoom(roomType string, roomId int32)
	GetRoomList(roomType string) map[int32]*Room
	GetRoom(roomId int32, roomType string) *Room
}

var RoomManagerInstance *RoomManager
var once sync.Once

func NewRoomManager() IRoomManager {
	once.Do(func() {
		RoomManagerInstance = &RoomManager{
			rooms: make(map[string]map[int32]*Room),
			Lock:  new(sync.RWMutex),
			ctx:   context.Background(),
		}
	})
	return RoomManagerInstance
}

func (rm *RoomManager) GetRoom(roomId int32, roomType string) *Room {
	return rm.GetRoomList(roomType)[roomId]
}

func (rm *RoomManager) CreateRoom(roomName, roomType string) *Room {
	rm.Lock.Lock()
	defer rm.Lock.Unlock()
	// ctx, cancel := context.WithCancel(rm.ctx)
	room := NewRoom(roomName, roomType, rm.ctx)
	rm.rooms[roomType][room.Id] = room
	go room.Run()
	return room
}

func (rm *RoomManager) DestroyRoom(roomType string, roomId int32) {
	rm.Lock.Lock()
	defer rm.Lock.Unlock()
	delete(rm.rooms[roomType], roomId)
}

func (rm *RoomManager) GetRoomList(roomType string) map[int32]*Room {
	rm.Lock.RLock()
	defer rm.Lock.RUnlock()
	return rm.rooms[roomType]
}
