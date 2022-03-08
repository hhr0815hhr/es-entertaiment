package proto

import (
	"context"
	"es-entertainment/core/codec"
	"es-entertainment/lib/send"
	"es-entertainment/module/game"
	"es-entertainment/protos"

	"github.com/gorilla/websocket"
)

type CreateRoom struct {
	protos.C2S_CreateRoom
}

func (l *CreateRoom) Handle(ctx context.Context, data []byte) {
	player_id := ctx.Value("value").(map[string]interface{})["player"].(int)
	pp := &protos.C2S_CreateRoom{}
	codec.Instance().Decode(data, pp)
	r := game.LobbyInstance.RoomManager.CreateRoom(pp.RoomName, pp.RoomType)
	err := r.Join(player_id)
	if err != nil {
		send.SendTo(ctx.Value("value").(map[string]interface{})["conn"].(*websocket.Conn), err.Error())
	}
}

type Ready struct{}

func (r *Ready) Handle(ctx context.Context, data []byte) {
	// player_id := ctx.Value("value").(map[string]interface{})["player"].(int)
	// room_id := ctx.Value("value").(map[string]interface{})["room"].(int)
}
