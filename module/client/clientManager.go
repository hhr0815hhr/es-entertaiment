package client

import (
	"errors"
	"sync"
)

type ClientManager struct {
	clients map[int64]*Client
	lock    *sync.Mutex
}

var (
	ClientManagerIns *ClientManager
)

func init() {
	ClientManagerIns = &ClientManager{
		clients: make(map[int64]*Client),
		lock:    &sync.Mutex{},
	}
}

func (cm *ClientManager) Create(id int64, conn CliConn) error {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	if _, ok := cm.clients[id]; ok {
		return errors.New("client Id already exists")
	}
	cli := &Client{
		Id:   id,
		Conn: conn,
	}
	cm.clients[id] = cli
	return nil
}
