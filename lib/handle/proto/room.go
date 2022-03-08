package proto

import (
	"context"
	"es-entertainment/core/codec"
	"es-entertainment/core/log"
	"es-entertainment/lib/send"
	"es-entertainment/module/game"
	"es-entertainment/protos"
)

type RoomList struct {
}

func (l *RoomList) Handle(ctx context.Context, data []byte) {
	pp := &protos.C2S_RoomList{}
	codec.Instance().Decode(data, pp)
	list := game.LobbyInstance.RoomManager.GetRoomList(pp.RoomType)
	roomList := make([]*protos.RoomInfo, 0)
	for x := range list {
		roomList = append(roomList, &protos.RoomInfo{
			RoomId:        list[x].Id,
			RoomPlayerNum: int32(len(list[x].Players)),
			RoomStatus:    list[x].State,
		})
	}
	msg := &protos.S2C_RoomList{
		RoomList: roomList,
	}
	b, err := codec.Instance().Encode(msg)
	if err != nil {
		log.Errorf("encode error: %s", err)
	} else {
		send.SendToUid(ctx.Value("value").(map[string]interface{})["playerId"].(int64), b, int32(protos.CmdType_CMD_S2C_RoomList))
	}

}

type CreateRoom struct {
	// protos.C2S_CreateRoom
}

func (l *CreateRoom) Handle(ctx context.Context, data []byte) {
	player_id := ctx.Value("value").(map[string]interface{})["playerId"].(int64)
	pp := &protos.C2S_CreateRoom{}
	codec.Instance().Decode(data, pp)
	r := game.LobbyInstance.RoomManager.CreateRoom(pp.RoomName, pp.RoomType)
	err := r.Join(player_id)
	ret := &protos.S2C_CreateRoom{}
	if err != nil {
		log.Errorf("join room error: %s", err)
		ret.Code = 1
	} else {
		ret.Code = 0
	}
	var b []byte
	b, err = codec.Instance().Encode(ret)
	if err != nil {
		log.Errorf("encode error: %s", err)
	} else {
		send.SendToUid(player_id, b, int32(protos.CmdType_CMD_S2C_CreateRoom))
	}
}

type EnterRoom struct{}

func (r *EnterRoom) Handle(ctx context.Context, data []byte) {
	pp := &protos.C2S_EnterRoom{}
	codec.Instance().Decode(data, pp)
	ret := &protos.S2C_EnterRoom{}
	player_id := ctx.Value("value").(map[string]interface{})["playerId"].(int64)
	if room, ok := game.LobbyInstance.GetRooms(pp.RoomType)[pp.RoomId]; ok {
		err := room.Join(player_id)
		if err != nil {
			log.Errorf("join room error: %s", err)
			ret.Code = 1
		} else {
			ret.Code = 0
		}
	} else {
		ret.Code = 1
	}
	b, err := codec.Instance().Encode(ret)
	if err != nil {
		log.Errorf("encode error: %s", err)
	} else {
		send.SendToUid(player_id, b, int32(protos.CmdType_CMD_S2C_EnterRoom))
	}
}

type LeaveRoom struct{}

func (r *LeaveRoom) Handle(ctx context.Context, data []byte) {
	pp := &protos.C2S_LeaveRoom{}
	codec.Instance().Decode(data, pp)

}

type Ready struct{}

func (r *Ready) Handle(ctx context.Context, data []byte) {
	// player_id := ctx.Value("value").(map[string]interface{})["player"].(int)
	// room_id := ctx.Value("value").(map[string]interface{})["room"].(int)
}
