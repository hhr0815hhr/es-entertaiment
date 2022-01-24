package main

import (
	cs "es-entertainment/core/server"
	"es-entertainment/module/game"
	"es-entertainment/module/server"
	"flag"
)

var (
	//flag参数
	serverType string
)

func init() {
	flag.StringVar(&serverType, "s", cs.Server_Type_WS, "server type: tcp, udp, ws")
}

func main() {
	flag.Parse()
	initModules()
	err := server.Run(serverType)
	if err != nil {
		panic(err)
	}
}

func initModules() {
	game.InitLobby()
	// player.InitPlayer()
}
