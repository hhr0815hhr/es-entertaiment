package handle

import (
	"context"
	"es-entertainment/core/pack"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandle(c *gin.Context) {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		defer conn.Close()
		for {
			ctxMap := make(map[string]interface{})
			ctxMap["conn"] = conn
			ctx := context.WithValue(context.Background(), "value", ctxMap)
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			// fmt.Println("原数据：", string(message))
			//解包
			data, err := pack.Instance().Unpack(message)
			// fmt.Println("解包数据：", data)
			if err != nil {
				fmt.Println(err)
				continue
			}

			dispatch(ctx, data)
			// err = conn.WriteMessage(websocket.TextMessage, message)
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }
		}
	}()
}
