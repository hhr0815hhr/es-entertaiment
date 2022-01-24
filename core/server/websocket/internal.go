package websocket

// "github.com/gorilla/websocket"
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"
)

var upGrader = ws.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serve(host string, port int) error {
	r := gin.Default()
	r.GET("/ws", wsHandle)
	err := r.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func wsHandle(c *gin.Context) {
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
			//
			fmt.Println(string(message))
			err = conn.WriteMessage(ws.TextMessage, message)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
}
