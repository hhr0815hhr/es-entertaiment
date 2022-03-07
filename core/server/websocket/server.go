package websocket

import "github.com/gin-gonic/gin"

type WsServer struct {
	WsHandle   func(c *gin.Context)
	HttpHandle func(c *gin.Context)
}

func NewServer(id int64) *WsServer {
	return &WsServer{
		// Id: id,
	}
}

func (ws *WsServer) BeforeRun() {

}

func (ws *WsServer) SetHandle(fWs interface{}, fHttp interface{}) {
	ws.WsHandle = fWs.(func(c *gin.Context))
	ws.HttpHandle = fHttp.(func(c *gin.Context))
}

func (ws *WsServer) Run(host string, port int) error {
	err := serve(host, port, ws.WsHandle, ws.HttpHandle)
	if err != nil {
		return err
	}
	return nil
}

func (ws *WsServer) Stop() {

}
