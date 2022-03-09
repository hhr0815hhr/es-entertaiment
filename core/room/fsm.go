package room

import (
	"github.com/looplab/fsm"
)

//room logic can be treated as a fsm

type F struct {
	Name string
	FSM  *fsm.FSM
}

func initFsm() {

}
