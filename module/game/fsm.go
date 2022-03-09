package game

import (
	"github.com/looplab/fsm"
)

type FsmManager struct {
	fsmMap map[string]*fsm.FSM
}

func (fm *FsmManager) GetFsm(name string) *fsm.FSM {
	return fm.fsmMap[name]
}

func (fm *FsmManager) RegisterFsm(name string, f *fsm.FSM) {
	fm.fsmMap[name] = f
}


