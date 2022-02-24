package proto

import (
	"context"
	"es-entertainment/module/game"
)

type Login struct{}

func (l *Login) Handle(ctx context.Context, data interface{}) {
	var p = 1
	game.LobbyInstance.EnterLobby(p)
}
