package main

import (
	"fmt"
)

func ElephantFeed() string {
	return "Grass"
}

func MonkeyFeed() string {
	return "Banana"
}

func RabbitFeed() string {
	return "Carrot"
}

func main() {
  fmt.Println(AppName())

	fmt.Println(ElephantFeed())
	fmt.Println(MonkeyFeed())
	fmt.Println(RabbitFeed())
}

