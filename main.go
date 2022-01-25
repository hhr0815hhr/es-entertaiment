package main

import (
	"es-entertainment/conf"
	"es-entertainment/core/database/mysql"
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
	cfg := conf.GetConf()
	initModules(cfg)
	err := server.Run(serverType, cfg)
	if err != nil {
		panic(err)
	}
}

func initModules(cfg conf.Config) {

	mysql.InitDB(cfg.Mysql["master"], cfg.Mysql["slave"])
	game.InitLobby()
	// player.InitPlayer()
}
