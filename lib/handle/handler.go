package handle

import (
	"es-entertainment/core/codec"
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
			//解码
			var decode DecodeStruct
			err = codec.Instance().Decode(data.([]byte), &decode)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("解码数据：", decode)
			dispatch(decode.Cmd, decode.Data)
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}()
}
