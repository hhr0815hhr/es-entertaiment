package player

import "github.com/gorilla/websocket"

type Player struct {
	Id       int64   	`gorm:"primary_key"`
	Username string		`gorm:"type:varchar(32);unique_index"`
	Password string		`gorm:"type:varchar(32)"`
	Nick     string		`gorm:"type:varchar(32)"`
	Icon     string		`gorm:"type:varchar(32)"`
	Coin     int64		`gorm:"type:bigint"`
	Level    int		`gorm:"type:int"`
	Sex      int		`gorm:"type:int"`
	Other    PlayerOther
}

type PlayerOther struct {
	Conn         *websocket.Conn
	Position     int // 当前所在场景 0：大厅  10001：房间号
	PositionType string
}

const (
	Position_Lobby    = iota //"lobby"
	Position_Cow             //"cow"
	Position_Landlord        //"landlord"
	Position_Gomoku          //"gomoku"
)
