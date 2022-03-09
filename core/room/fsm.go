package room

import (
	"github.com/looplab/fsm"
)

//room logic can be treated as a fsm

type Fsm struct {
	Name string
	FSM  *fsm.FSM
}

func initFsm(name string, f *fsm.FSM) *Fsm {
	return &Fsm{
		Name: name,
		FSM:  f,
	}
}


func getFsm(roomType string) *fsm.FSM {
	return &fsm.FSM{}
}