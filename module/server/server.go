package server

import (
	"es-entertainment/core/server"
)

func Run(serverType string) error {
	//get conf  host string, port int
	s, err := server.NewServer(1, server.Server_Type_WS)
	if err != nil {
		return nil
	}
	host := "0.0.0.0"
	port := 8567
	// s.BeforeRun()
	err = s.Run(host, port)
	return err
}
