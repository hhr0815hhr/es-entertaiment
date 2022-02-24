package send

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func SendTo(conn *websocket.Conn, data interface{}) {
	if conn == nil {
		return
	}
	err := conn.WriteJSON(data)
	if err != nil {
		fmt.Println(err)
	}
}

func MultiSend(conns []*websocket.Conn, data interface{}) {
	for _, conn := range conns {
		SendTo(conn, data)
	}
}
