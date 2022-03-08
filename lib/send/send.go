package send

import (
	"es-entertainment/common"
	"es-entertainment/core/log"
	"es-entertainment/lib/player"
	p "es-entertainment/module/data/player"

	"github.com/gorilla/websocket"
)

func SendToUid(uid int64, data []byte, cmd int32) {
	if a, ok := player.PlayerMap.Load(uid); ok {
		SendTo(a.(*p.Player).Other.Conn, data, cmd)
	}
}

func SendToUids(uids []int64, data []byte, cmd int32) {
	conns := make([]*websocket.Conn, 0)
	for _, uid := range uids {
		if a, ok := player.PlayerMap.Load(uid); ok {
			conns = append(conns, a.(*p.Player).Other.Conn)
		}
	}
	MultiSend(conns, data, cmd)
}

func SendTo(conn *websocket.Conn, data []byte, cmd int32) {
	if conn == nil {
		return
	}
	buf := common.IntToBytes(int(cmd))

	err := conn.WriteMessage(websocket.BinaryMessage, buf)
	if err != nil {
		log.Errorf("send error: %s", err)
	}
}

func MultiSend(conns []*websocket.Conn, data []byte, cmd int32) {
	for _, conn := range conns {
		SendTo(conn, data, cmd)
	}
}
