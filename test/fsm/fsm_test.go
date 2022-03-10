package fsm

import (
	"es-entertainment/module/game/cow"
	"testing"
)

func TestFsm(t *testing.T) {
	fsm := cow.NewFsm()
	fsm.Event("start")
	// fsm.Event("draw")
	// fsm.Event("compare")
	// fsm.Event("ready")
	
}
