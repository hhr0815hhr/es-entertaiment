package cow

import (
	"es-entertainment/common"
)

func CalcCow(cards []int32) (haveCow bool, cowType int32) {
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

func formatCards(cards []int32) ([]int32, bool) {
	isFlower := true
	cp := make([]int32, 5)
	copy(cp, cards)
	for i := 0; i < len(cp); i++ {
		cp[i] = int32(cp[i] / 10)
		if cp[i] >= 10 {
			cp[i] = 10
		} else {
			isFlower = false
		}
	}
	return cp, isFlower
}

func getLargestCard(cards []int32) int32 {
	var max int32 = 0
	for _, v := range cards {
		if v > max {
			max = v
		}
	}
	return max
}

func Compare(masterCards, playerCards []int32) int32 {
	if masterCards[0] == playerCards[0] {
		return 0
	}
	_, v1 := CalcCow(masterCards)
	_, v2 := CalcCow(playerCards)
	if v1 > v2 {
		return 1
	}
	if v1 < v2 {
		return 2
	}
	if getLargestCard(masterCards) > getLargestCard(playerCards) {
		return 1
	}
	return 2
}
