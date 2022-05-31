package main

import "fmt"

func main() {
	defer func(){
		fmt.Println("defer")
  }()
	s := "A"
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
		fallthrough
	case "C":
		s += "D"
		fallthrough
	default:
		s += "E"
	}
	fmt.Println(s)
	defer fmt.Println("last")
}
