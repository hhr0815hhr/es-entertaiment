package websocket

// "github.com/gorilla/websocket"
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func serve(host string, port int, wsHandle func(ctx *gin.Context), httpHandle func(ctx *gin.Context)) error {
	r := gin.New()
	r.Any("/", httpHandle)
	r.GET("/ws", wsHandle)
	err := r.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
