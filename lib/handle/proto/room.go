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
	list := game.LobbyInstance.GetRooms(pp.RoomType)
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
	b, _ := codec.Instance().Encode(msg)

	send.SendToUid(ctx.Value("value").(map[string]interface{})["playerId"].(int64), b, int32(protos.CmdType_CMD_S2C_RoomList))

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
		game.LobbyInstance.LeaveLobby(player_id)
	}

	b, _ := codec.Instance().Encode(ret)

	r.Cast(player_id, int32(protos.CmdType_CMD_S2C_CreateRoom), b)
}

type EnterRoom struct{}

func (r *EnterRoom) Handle(ctx context.Context, data []byte) {
	pp := &protos.C2S_EnterRoom{}
	codec.Instance().Decode(data, pp)
	player_id := ctx.Value("value").(map[string]interface{})["playerId"].(int64)
	ret := &protos.S2C_EnterRoom{
		Code:     0,
		PlayerId: player_id,
	}

	room := game.LobbyInstance.RoomManager.GetRoom(pp.RoomId, pp.RoomType)

	err := room.Join(player_id)
	if err != nil {
		log.Errorf("join room error: %s", err)
		ret.Code = 1
	} else {
		ret.Code = 0
		game.LobbyInstance.LeaveLobby(player_id)
	}

	b, _ := codec.Instance().Encode(ret)
	if ret.Code == 0 {
		room.Broadcast(int32(protos.CmdType_CMD_S2C_EnterRoom), b)
	} else {
		send.SendToUid(player_id, b, int32(protos.CmdType_CMD_S2C_EnterRoom))
	}
}

type LeaveRoom struct{}

func (r *LeaveRoom) Handle(ctx context.Context, data []byte) {
	pp := &protos.C2S_LeaveRoom{}
	codec.Instance().Decode(data, pp)
	player_id := ctx.Value("value").(map[string]interface{})["playerId"].(int64)
	room := game.LobbyInstance.RoomManager.GetRoom(pp.RoomId, pp.RoomType)
	err := room.Leave(player_id)
	ret := &protos.S2C_LeaveRoom{
		Code:     0,
		PlayerId: player_id,
	}
	if err != nil {
		ret.Code = 1
	} else {
		game.LobbyInstance.EnterLobby(player_id)
	}
	b, _ := codec.Instance().Encode(ret)

	send.SendToUid(player_id, b, int32(protos.CmdType_CMD_S2C_LeaveRoom))
	room.Broadcast(int32(protos.CmdType_CMD_S2C_LeaveRoom), b)
}

type Ready struct{}

func (r *Ready) Handle(ctx context.Context, data []byte) {
	pp := &protos.C2S_Ready{}
	codec.Instance().Decode(data, pp)
	player_id := ctx.Value("value").(map[string]interface{})["playerId"].(int64)
	room := game.LobbyInstance.RoomManager.GetRoom(pp.RoomId, pp.RoomType)
	err := room.Ready(player_id, int(pp.Ready))
	ret := &protos.S2C_Ready{
		Code:     0,
		PlayerId: player_id,
	}
	if err != nil {
		ret.Code = 1
	}

	b, _ := codec.Instance().Encode(ret)

	room.Broadcast(int32(protos.CmdType_CMD_S2C_Ready), b)
}
