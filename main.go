package main

import (
	"es-entertainment/conf"
	"es-entertainment/core/database/mysql"
	"es-entertainment/core/log"
	cs "es-entertainment/core/server"
	"es-entertainment/lib/player"
	"es-entertainment/module/game"
	"es-entertainment/module/server"
	"flag"
	"os"
	"os/signal"
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
	go func() {
		err := server.Run(serverType, cfg)
		if err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	log.Info("server stop...")
}

func initModules(cfg conf.Config) {
	log.InitLogger()
	mysql.InitDB(cfg.Mysql["master"], cfg.Mysql["slave"])
	game.InitLobby()
	game.InitFsm()
	player.InitPlayer()
}
