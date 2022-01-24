package websocket

import "github.com/gin-gonic/gin"

type WsServer struct {
	WsHandle func(c *gin.Context)
}

func NewServer(id int64) *WsServer {
	return &WsServer{
		// Id: id,
	}
}

func (ws *WsServer) BeforeRun() {

}

func (ws *WsServer) SetHandle(f interface{}) {
	ws.WsHandle = f.(func(c *gin.Context))
}

func (ws *WsServer) Run(host string, port int) error {
	err := serve(host, port, ws.WsHandle)
	if err != nil {
		return err
	}
	return nil
}

func (ws *WsServer) Stop() {

}
