package server

import (
	"es-entertainment/core/server"
	"es-entertainment/lib/handle"
)

func Run(serverType string) error {
	//get conf  host string, port int
	s, err := server.NewServer(1, server.Server_Type_WS)
	s.SetHandle(handle.WsHandle)
	if err != nil {
		return nil
	}
	host := "0.0.0.0"
	port := 8567
	// s.BeforeRun()
	err = s.Run(host, port)
	return err
}
