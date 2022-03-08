package player

import (
	"es-entertainment/core/log"
	"sync"
)

var PlayerMap *PlayerList

func InitPlayer() {
	PlayerMap = new(PlayerList)
	PlayerMap.Map = &sync.Map{}
	log.Info("player init success...")
}
