package main

import "fmt"

func operation1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func operation2(s []int, c chan int) {
	sum := 100
	for _, v := range s {
		sum *= v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go operation1(s, c)
	go operation2(s, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
