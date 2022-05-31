package main

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("Here")
	fmt.Println("NumCPU: %d\n", runtime.NumCPU())
	fmt.Println("NumGoRoutine: %d\n", runtime.NumGoroutine())
	fmt.Println("Version: %d\n", runtime.Version())
}
