package main

import "fmt"

func sub() {
	for {
		fmt.Println("sub loop")
	}
}

func goFunk() {
	go sub()
	for {
		fmt.Println("main loop")
	}
}
