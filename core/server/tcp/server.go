package tcp

type TcpServer struct{}

func NewServer(id int64) *TcpServer {
	return &TcpServer{
		// Id: id,
	}
}

func (ts *TcpServer) SetHandle(f interface{}, httpF interface{}) {

}

func (ts *TcpServer) BeforeRun() {

}

func (ts *TcpServer) Run(host string, port int) error {
	return nil
}

func (ts *TcpServer) Stop() {

}
