package common

func InSlice(a interface{}, list []interface{}) bool {
	// reflect.TypeOf(a)
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
