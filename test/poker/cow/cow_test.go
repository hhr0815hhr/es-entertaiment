package cow

import (
	"es-entertainment/common"
	"es-entertainment/lib/poker"
	c "es-entertainment/module/game/cow"
	"testing"
)

func TestCalcCow(t *testing.T) {
	poker.InitCards()
	common.ShuffleSlice(poker.AllCards)
	cards := poker.AllCards[:5]
	haveCow, cowType := c.CalcCow(cards)
	t.Logf("cards %v,CalcCow got:%v %d", cards, haveCow, cowType)

}
