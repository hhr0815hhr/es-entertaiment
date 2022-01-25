package websocket

// "github.com/gorilla/websocket"
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func serve(host string, port int, wsHandle func(ctx *gin.Context)) error {
	r := gin.Default()
	r.GET("/", httpHandle)
	r.GET("/ws", wsHandle)
	err := r.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func httpHandle(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/favicon.ico" {
		ctx.Abort()
		return
	}
	fmt.Println(ctx.Request.URL.Path)
	ctx.String(200, "hello")
}
