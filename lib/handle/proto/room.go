package proto

import (
	"context"
	"es-entertainment/lib/send"
	"es-entertainment/module/game"

	"github.com/gorilla/websocket"
)

type CreateRoom struct {
	RoomType string `json:"roomType"`
	RoomName string `json:"roomName"`
}

func (l *CreateRoom) Handle(ctx context.Context, data interface{}) {
	player_id := ctx.Value("value").(map[string]interface{})["player"].(int)
	r := game.LobbyInstance.RoomManager.CreateRoom(data.(CreateRoom).RoomName, data.(CreateRoom).RoomType)
	err := r.Join(player_id)
	if err != nil {
		send.SendTo(ctx.Value("value").(map[string]interface{})["conn"].(*websocket.Conn), err.Error())
	}
}

type Ready struct{}

func (r *Ready) Handle(ctx context.Context, data interface{}) {
	player_id := ctx.Value("value").(map[string]interface{})["player"].(int)
	room_id := ctx.Value("value").(map[string]interface{})["room"].(int)
}
