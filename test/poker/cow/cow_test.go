package cow

import (
	"es-entertainment/common"
	"es-entertainment/lib/poker"
	c "es-entertainment/module/game/cow"
	"testing"
)

func TestCalcCow(t *testing.T) {
	pokerCopy := poker.InitCards()
	common.ShuffleSlice(pokerCopy)
	cards := pokerCopy[:5]
	haveCow, cowType := c.CalcCow(cards)
	t.Logf("cards %v,CalcCow got:%v %d", cards, haveCow, cowType)

}
