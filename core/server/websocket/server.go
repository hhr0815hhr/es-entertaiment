package websocket

type WsServer struct{}

func NewServer(id int64) *WsServer {
	return &WsServer{
		// Id: id,
	}
}

func (ws *WsServer) BeforeRun() {

}

func (ws *WsServer) Run(host string, port int) error {
	err := serve(host, port)
	if err != nil {
		return err
	}
	return nil
}

func (ws *WsServer) Stop() {

}
