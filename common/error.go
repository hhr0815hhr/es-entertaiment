package common

import "fmt"

func CatchPanic(f func()) (err interface{}) {
	defer func() {
		if err = recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	f()
	return
}

func RunPanicless(f func()) (panicless bool) {
	defer func() {
		err := recover()
		panicless = err == nil
		if err != nil {
			fmt.Printf("%s panic: %s", f, err)
		}
	}()
	f()
	return
}

func RunNoPanic(f func()) {
	for !RunPanicless(f) {
	}
}
