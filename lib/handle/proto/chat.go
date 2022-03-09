package proto

import (
	"context"
	"es-entertainment/core/codec"
	c "es-entertainment/lib/chat"
	"es-entertainment/module/game"
	"es-entertainment/protos"
)

type Chat struct{}

// const (
// 	Channel_Lobby = 1
// )

func getChatChannel(channel int32) chan *c.Chat {
	switch channel {
	case int32(protos.ChatChannel_Lobby):
		return game.LobbyInstance.LobbyChatChannel
	default:
		return nil
	}
}

func (ct *Chat) Handle(ctx context.Context, data []byte) {
	// var p = 1
	player_id := ctx.Value("value").(map[string]interface{})["playerId"].(int64)
	pp := &protos.C2S_Chat{}
	codec.Instance().Decode(data, pp)
	ch := getChatChannel(pp.Channel)
	msg := &c.Chat{
		From: player_id,
		Msg:  pp.Msg,
		Time: pp.Time,
	}
	ch <- msg
}
