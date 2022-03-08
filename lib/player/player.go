package player

import (
	"es-entertainment/core/log"
	"es-entertainment/module/data"
	p "es-entertainment/module/data/player"
	"sync"
)

var PlayerMap *PlayerList

func InitPlayer() {
	PlayerMap = new(PlayerList)
	PlayerMap.Map = &sync.Map{}
	log.Info("player init success...")
}

func GetPlayerByUsername(username string) (p *p.Player) {
	data.ReadDb().Table("player").Where("username = ?", username).First(p)
	return
}
