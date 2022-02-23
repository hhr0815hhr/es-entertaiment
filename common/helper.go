package common

import (
	"math/rand"
	"time"
)

func InSlice(a interface{}, list []interface{}) bool {
	// reflect.TypeOf(a)
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SumSlice(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func ShuffleSlice(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
