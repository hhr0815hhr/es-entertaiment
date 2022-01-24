package udp

type UdpServer struct{}

func NewServer(id int64) *UdpServer {
	return &UdpServer{
		// Id: id,
	}
}

func (us *UdpServer) SetHandle(f interface{}) {

}

func (us *UdpServer) BeforeRun() {

}

func (us *UdpServer) Run(host string, port int) error {
	return nil
}

func (us *UdpServer) Stop() {

}
