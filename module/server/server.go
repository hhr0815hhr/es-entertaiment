package server

import (
	"es-entertainment/conf"
	"es-entertainment/core/server"
	"es-entertainment/lib/handle"
)

func Run(serverType string, cfg conf.Config) error {
	//get conf  host string, port int
	s, err := server.NewServer(1, server.Server_Type_WS)
	s.SetHandle(handle.WsHandle)
	if err != nil {
		return nil
	}
	// s.BeforeRun()
	err = s.Run(cfg.Server.Host, cfg.Server.Port)
	return err
}
