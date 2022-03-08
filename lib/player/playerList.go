package player

import (
	p "es-entertainment/module/data/player"
	"sync"
)

type PlayerList struct {
	*sync.Map
}

func LoadPlayer(p *p.Player) {
	PlayerMap.Store(p.Id, p)
}
