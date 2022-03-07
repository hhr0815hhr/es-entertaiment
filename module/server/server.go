package server

import (
	"es-entertainment/conf"
	"es-entertainment/core/server"
	"es-entertainment/lib/handle"
)

func Run(serverType string, cfg conf.Config) error {
	//get conf  host string, port int
	s, err := server.NewServer(1, serverType)
	if err != nil {
		return err
	}
	if serverType == "ws" {
		s.SetHandle(handle.WsHandle, handle.HttpHandle)
	}
	// s.BeforeRun()
	err = s.Run(cfg.Server.Host, cfg.Server.Port)
	return err
}
