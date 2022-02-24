package common

import (
	"errors"
	"fmt"
)

func CatchPanic(f func()) (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic:", p)
			err = errors.New(p.(string))
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
			fmt.Printf("%T panic: %s", f, err)
		}
	}()
	f()
	return
}

func RunNoPanic(f func()) {
	for !RunPanicless(f) {
	}
}
