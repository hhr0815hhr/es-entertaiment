package websocket

// "github.com/gorilla/websocket"
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func serve(host string, port int, handle func(ctx *gin.Context)) error {
	r := gin.Default()
	r.GET("/ws", handle)
	err := r.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
