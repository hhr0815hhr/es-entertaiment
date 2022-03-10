package room

import (
	"es-entertainment/core/codec"
	"es-entertainment/protos"
	"fmt"
	"reflect"
	"time"

	"github.com/looplab/fsm"
)

type RoomFsm fsm.FSM

func (f *RoomFsm) Event(event string, args ...interface{}) error {
	return reflect.ValueOf(f).Interface().(*fsm.FSM).Event(event, args...)
}

func getRoomFsm(roomType string) *RoomFsm {
	funcName := fmt.Sprintf("%sFsm", roomType)
	f := reflect.ValueOf(funcName).Interface().(func() *fsm.FSM)()
	return reflect.ValueOf(f).Interface().(*RoomFsm)
	// switch roomType {
	// case "cow":
	// 	f = cowFsm()
	// default:
	// 	f = nil
	// }
	// return
}

func cowFsm() *RoomFsm {
	f := fsm.NewFSM(
		"ready",
		fsm.Events{
			{Name: "start", Src: []string{"ready"}, Dst: "start"},
			{Name: "draw", Src: []string{"start"}, Dst: "draw"},
			{Name: "compare", Src: []string{"draw"}, Dst: "compare"},
			{Name: "ready", Src: []string{"compare"}, Dst: "ready"},
		},
		fsm.Callbacks{
			// "enter_state": func(e *fsm.Event) {
			// 	fmt.Printf("触发事件%s", e.Event)
			// },
			"enter_start":   startCowEvent,
			"enter_draw":    drawCowEvent,
			"enter_compare": compareCowEvent,
			"enter_ready":   readyCowEvent,
		},
	)
	return reflect.ValueOf(f).Interface().(*RoomFsm)
}

//event只会由玩家操作和定时器两项触发
//触发event时，重置ticker

func startCowEvent(e *fsm.Event) {
	ro := e.Args[0].(*Room)
	ro.EventTimer.Stop()
	ticker := ro.Tickers["start"]
	fmt.Printf("进入状态%s", e.Dst)

	ro.State = 1
	ret := &protos.S2C_Cow_Start{
		CountDown: int32(ticker.Time),
		State:     ro.State,
	}
	b, _ := codec.Instance().Encode(ret)
	ro.Broadcast(int32(protos.CmdType_CMD_S2C_Cow_Start), b)
	ro.EventTimer = time.AfterFunc(ticker.Time, func() { ro.F.Event(ticker.Event, ro) })
}

func drawCowEvent(e *fsm.Event) {
	ro := e.Args[0].(*Room)
	ro.EventTimer.Stop()
	ticker := ro.Tickers["draw"]
	fmt.Printf("进入状态%s", e.Dst)

	ro.State = 2
	ret := &protos.S2C_Cow_Draw{
		CountDown: int32(ticker.Time),
		State:     ro.State,
	}
	b, _ := codec.Instance().Encode(ret)
	ro.Broadcast(int32(protos.CmdType_CMD_S2C_Cow_Draw), b)
	ro.EventTimer = time.AfterFunc(ticker.Time, func() { ro.F.Event(ticker.Event, ro) })
}

func compareCowEvent(e *fsm.Event) {
	ro := e.Args[0].(*Room)
	ro.EventTimer.Stop()
	ticker := ro.Tickers["compare"]
	fmt.Printf("进入状态%s", e.Dst)

	ro.State = 3
	ret := &protos.S2C_Cow_Compare{
		CountDown: int32(ticker.Time),
		State:     ro.State,
	}
	b, _ := codec.Instance().Encode(ret)
	ro.Broadcast(int32(protos.CmdType_CMD_S2C_Cow_Compare), b)
	ro.EventTimer = time.AfterFunc(ticker.Time, func() { ro.F.Event(ticker.Event, ro) })
}

func readyCowEvent(e *fsm.Event) {
	ro := e.Args[0].(*Room)
	ro.EventTimer.Stop()

	fmt.Printf("进入状态%s", e.Dst)

	ro.State = 0
	ret := &protos.S2C_Cow_Ready{
		State:     ro.State,
	}
	b, _ := codec.Instance().Encode(ret)
	ro.Broadcast(int32(protos.CmdType_CMD_S2C_Cow_Ready), b)
}
