package proto

import (
	"context"
	"es-entertainment/common"
	"es-entertainment/core/codec"
	"es-entertainment/core/log"
	"es-entertainment/lib/player"
	"es-entertainment/lib/send"
	"es-entertainment/module/game"
	"es-entertainment/protos"

	"github.com/gorilla/websocket"
)

type Login struct {
	// protos.C2S_Login
}

func (l *Login) Handle(ctx context.Context, data []byte) {
	// var p = 1
	pp := &protos.C2S_Login{}
	codec.Instance().Decode(data, pp)

	p := player.GetPlayerByUsername(pp.Username)

	ret := &protos.S2C_Login{}
	if p != nil && p.Password == string(common.Encrypt([]byte(pp.Password))) {
		//playerInfo存入playerList
		p.Other.Conn = ctx.Value("value").(map[string]interface{})["conn"].(*websocket.Conn)
		delete(ctx.Value("value").(map[string]interface{}), "conn") //conn存入player结构体 不存在上下文中
		player.LoadPlayer(p)

		//playerId 存入 ctx
		ctx.Value("value").(map[string]interface{})["playerId"] = p.Id
		game.LobbyInstance.EnterLobby(p.Nick)
		ret.Code = 0
	} else {
		ret.Code = 1
	}
	b, err := codec.Instance().Encode(ret)
	if err != nil {
		log.Errorf("encode error: %s", err)
	}
	send.SendTo(p.Other.Conn, b)
}
