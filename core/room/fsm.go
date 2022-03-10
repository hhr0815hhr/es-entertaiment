package room

import (
	"github.com/looplab/fsm"
)

//room logic can be treated as a fsm
type FsmManager struct {
	FsmMap map[string]*fsm.FSM
}

var FsmM *FsmManager

func (fm *FsmManager) GetFsm(name string) *fsm.FSM {
	f, ok := fm.FsmMap[name]
	if !ok {
		return nil
	}
	return f
}

func (fm *FsmManager) RegisterFsm(name string, f *fsm.FSM) {
	fm.FsmMap[name] = f
}

// func initFsm(name string, f *fsm.FSM) *Fsm {
// 	return &Fsm{
// 		Name: name,
// 		FSM:  f,
// 	}
// }

// func getFsm(roomType string) *fsm.FSM {
// 	return &fsm.FSM{}
// }
