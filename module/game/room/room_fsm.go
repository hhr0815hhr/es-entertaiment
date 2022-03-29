package room

import (
	"es-entertainment/common"
	"es-entertainment/core/codec"
	"es-entertainment/lib/poker"
	"es-entertainment/module/game/cow"
	"es-entertainment/protos"
	"fmt"
	"math/rand"
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

const (
	CowStateReady = iota
	CowStateStart
	CowStateMaster
	CowStateDraw
	CowStateCompare
)

func cowFsm() *RoomFsm {
	f := fsm.NewFSM(
		"ready",
		fsm.Events{
			{Name: "start", Src: []string{"ready"}, Dst: "start"},
			{Name: "master", Src: []string{"start"}, Dst: "master"},
			{Name: "draw", Src: []string{"start", "master"}, Dst: "draw"},
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

	var ret interface{}
	//先判断有无zj
	if ro.MasterPos == -1 {
		//rand master
		ro.State = CowStateMaster
		ret = &protos.S2C_Cow_Master{
			CountDown: int32(ticker.Time),
			State:     ro.State,
		}
	} else {
		ro.State = CowStateStart
		ret = &protos.S2C_Cow_Start{
			CountDown: int32(ticker.Time),
			State:     ro.State,
		}

	}
	b, _ := codec.Instance().Encode(ret)
	ro.Broadcast(int32(protos.CmdType_CMD_S2C_Cow_Start), b)
	ro.EventTimer = time.AfterFunc(ticker.Time, func() { ro.F.Event(ticker.Event, ro) })
}

func masterCowEvent(e *fsm.Event) {
	ro := e.Args[0].(*Room)
	ro.EventTimer.Stop()
	ticker := ro.Tickers["master"]
	fmt.Printf("进入状态%s", e.Dst)
	ro.State = CowStateMaster
	rand.Seed(time.Now().UnixNano())
	ro.MasterPos = rand.Int31n(int32(len(ro.Players)))
	ret := &protos.S2C_Cow_Master{
		CountDown: int32(ticker.Time),
		State:     ro.State,
		MasterPos: ro.MasterPos,
	}
	b, _ := codec.Instance().Encode(ret)
	ro.Broadcast(int32(protos.CmdType_CMD_S2C_Cow_Master), b)
	ro.EventTimer = time.AfterFunc(ticker.Time, func() { ro.F.Event(ticker.Event, ro) })
}

func drawCowEvent(e *fsm.Event) {
	ro := e.Args[0].(*Room)
	ro.EventTimer.Stop()
	ticker := ro.Tickers["draw"]
	fmt.Printf("进入状态%s", e.Dst)

	//draw cards
	cards := poker.InitCards()
	common.ShuffleSlice(cards)

	num := len(ro.Players)
	var index int
	CardsMap := make([]*protos.Cards, 0, num)
	for i := 0; i < num; i++ {
		index = i * 5
		copy(ro.Players[i].PlayerCards, cards[index:index+5])
		CardsMap = append(CardsMap, &protos.Cards{
			Card: ro.Players[i].PlayerCards,
		})
	}

	ro.State = CowStateDraw
	ret := &protos.S2C_Cow_Draw{
		CountDown: int32(ticker.Time),
		State:     ro.State,
		Cards:     CardsMap,
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

	//比牌
	master := ro.Players[ro.MasterPos]
	CowTypes := make([]int32, 0, len(ro.Players))
	Result := make([]int32, 0, len(ro.Players))
	var change int32 = -1
	for k, v := range ro.Players {
		_, t := cow.CalcCow(v.PlayerCards)
		if k == int(ro.MasterPos) && t == cow.Type_Cow0 {
			change = 0
		}
		CowTypes = append(CowTypes, t)
		Result = append(Result, cow.Compare(master.PlayerCards, v.PlayerCards))
	}
	if change >= 0 {
		ro.MasterPos += 1
		if ro.MasterPos >= int32(len(ro.Players)) {
			ro.MasterPos = 0
		}
		change = ro.MasterPos
	}
	ro.State = CowStateCompare
	ret := &protos.S2C_Cow_Compare{
		CountDown: int32(ticker.Time),
		State:     ro.State,
		CowType:   CowTypes,
		Result:    Result,
		Change:    change,
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
		State: ro.State,
	}
	b, _ := codec.Instance().Encode(ret)
	ro.Broadcast(int32(protos.CmdType_CMD_S2C_Cow_Ready), b)
}
