package proto

import "es-entertainment/module/game"

type Login struct{}

func (l *Login) Handle(data interface{}) {
	var p = 1
	game.LobbyInstance.EnterLobby(p)
}
