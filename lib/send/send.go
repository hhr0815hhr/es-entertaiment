package send

import (
	"es-entertainment/lib/player"
	p "es-entertainment/module/data/player"
	"fmt"

	"github.com/gorilla/websocket"
)

func SendToUid(uid int64, data []byte) {
	if a, ok := player.PlayerMap.Load(uid); ok {
		SendTo(a.(*p.Player).Other.Conn, data)
	}
}

func SendToUids(uids []int64, data []byte) {
	conns := make([]*websocket.Conn, 0)
	for _, uid := range uids {
		if a, ok := player.PlayerMap.Load(uid); ok {
			conns = append(conns, a.(*p.Player).Other.Conn)
		}
	}
	MultiSend(conns, data)
}

func SendTo(conn *websocket.Conn, data []byte) {
	if conn == nil {
		return
	}
	err := conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		fmt.Println(err)
	}
}

func MultiSend(conns []*websocket.Conn, data []byte) {
	for _, conn := range conns {
		SendTo(conn, data)
	}
}
