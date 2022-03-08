package handle

import (
	"context"
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

type CtxKey string

func WsHandle(c *gin.Context) {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	key := CtxKey("value")
	go func() {
		defer conn.Close()
		for {
			ctxMap := make(map[string]interface{})
			ctxMap["conn"] = conn //没必要把conn放到ctx里面 登录之后将conn转移至player结构体中
			ctx := context.WithValue(context.Background(), key, ctxMap)
			_, message, err := conn.ReadMessage()

			if err != nil {
				fmt.Println(err)
				return
			}
			// ws不用拆包分包
			// fmt.Println("原数据：", string(message))
			//解包
			// data, err := pack.Instance().Unpack(message)
			// fmt.Println("解包数据：", data)
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }
			dispatch(ctx, message)
		}
	}()
}

func HttpHandle(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/favicon.ico" {
		ctx.Abort()
		return
	}
	if ctx.Request.Method == "GET" {
		fmt.Println(ctx.Request.URL.Path)
		ctx.String(200, "hello")
	}
	if ctx.Request.Method == "POST" {
		//login
		ctx.String(200, "hello")
	}

}
