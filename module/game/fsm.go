package game

import (
	"es-entertainment/core/log"
	"es-entertainment/core/room"
	"es-entertainment/module/game/cow"

	"github.com/looplab/fsm"
)

func InitFsm() {
	room.FsmM = &room.FsmManager{
		FsmMap: make(map[string]*fsm.FSM),
	}
	room.FsmM.RegisterFsm("cow", cow.NewFsm())

	log.Info("init fsm success...")
}
