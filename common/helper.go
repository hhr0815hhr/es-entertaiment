package common

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
)

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("obj not in the target")
}

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
