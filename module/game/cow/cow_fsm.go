package cow

import (
	"fmt"

	"github.com/looplab/fsm"
)

func NewFsm() (f *fsm.FSM) {
	f = fsm.NewFSM(
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
			"enter_start":   startEvent,
			"enter_draw":    drawEvent,
			"enter_compare": compareEvent,
			"enter_ready":   readyEvent,
		},
	)
	return f
}

func startEvent(e *fsm.Event) {
	fmt.Printf("进入状态%s", e.Dst)
}

func drawEvent(e *fsm.Event) {
	fmt.Printf("进入状态%s", e.Dst)
}

func compareEvent(e *fsm.Event) {
	fmt.Printf("进入状态%s", e.Dst)
}

func readyEvent(e *fsm.Event) {
	fmt.Printf("进入状态%s", e.Dst)
}
