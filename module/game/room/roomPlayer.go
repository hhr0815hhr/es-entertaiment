package room

import (
	"es-entertainment/lib/player"
	p "es-entertainment/module/data/player"
)

type RoomPlayer struct {
	PlayerId     int64
	PlayerNick   string
	PlayerIcon   string
	PlayerSex    int
	PlayerLevel  int
	PlayerCoin   int64
	PlayerStatus int
	PlayerCards  []int32
}

func initRoomPlayer(playerId int64) *RoomPlayer {
	rp := &RoomPlayer{
		PlayerId: playerId,
	}
	if info, ok := player.PlayerMap.Load(playerId); ok {
		rp.PlayerCoin = info.(*p.Player).Coin
		rp.PlayerIcon = info.(*p.Player).Icon
		rp.PlayerNick = info.(*p.Player).Nick
		rp.PlayerSex = info.(*p.Player).Sex
		rp.PlayerLevel = info.(*p.Player).Level
		rp.PlayerStatus = 0
		rp.PlayerCards = make([]int32, 0, 5)
	}
	return rp
}
