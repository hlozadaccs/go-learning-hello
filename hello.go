package main

import (
	"fmt"

	greetings "github.hlozadaccs.com/go-learning-greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
