package cow

import (
	"es-entertainment/common"
)

func CalcCow(cards []int) (haveCow bool, cowType int) {
	// length := len(cards)
	cp, isFlower := formatCards(cards)
	if isFlower {
		return true, Type_CowsFlower
	}
	sum := common.SumSlice(cp)
	if sum <= 10 {
		return true, Type_CowsBit
	}
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				if (cp[i]+cp[j]+cp[k])%10 == 0 {
					haveCow = true
					break
				}
			}
		}
	}
	if haveCow {
		cowType = sum % 10
		if cowType == 0 {
			cowType = Type_Cows
		}
	}
	return
}

func formatCards(cards []int) ([]int, bool) {
	isFlower := true
	cp := make([]int, 5)
	copy(cp, cards)
	for i := 0; i < len(cp); i++ {
		cp[i] = int(cp[i] / 10)
		if cp[i] >= 10 {
			cp[i] = 10
		} else {
			isFlower = false
		}
	}
	return cp, isFlower
}
