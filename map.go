package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 5,
	}

	m["lemon"] = 3
	m["grape"] = 4

	for k, v := range m {
		fmt.Println(k, v)
	}
}
