package main

import (
	"fmt"
)

func returnFunc(a int) func() {
	return func() {
		fmt.Println("Input num is", a)
	}
}

func main() {
	f := returnFunc(10)
	f()
	returnFunc(10)()
}
