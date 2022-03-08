package proto

import (
	"context"
	"es-entertainment/core/codec"
	"es-entertainment/module/game"
	"es-entertainment/protos"
)

type Login struct {
	protos.C2S_Login
}

func (l *Login) Handle(ctx context.Context, data []byte) {
	var p = 1
	pp := &protos.C2S_Login{}
	codec.Instance().Decode(data, pp)
	//存入playerList

	game.LobbyInstance.EnterLobby(p)
}
