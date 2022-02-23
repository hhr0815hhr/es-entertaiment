package cow

import (
	"es-entertainment/common"
)

func CalcCow(cards []int) (haveCow bool, cowType int) {
	// length := len(cards)
	cards = formatCards(cards)
	sum := common.SumSlice(cards)
	if sum <= 10 {
		return true, 12
	}
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				if (cards[i]+cards[j]+cards[k])%10 == 0 {
					haveCow = true
					break
				}
			}
		}
	}
	if haveCow {

		cowType = sum % 10
		if cowType == 0 {
			cowType = 10
		}
	}
	return
}

func formatCards(cards []int) []int {
	for i := 0; i < len(cards); i++ {
		cards[i] = int(cards[i] / 10)
		if cards[i] > 10 {
			cards[i] = 10
		}
	}
	return cards
}
