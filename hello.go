package main

import (
	"fmt"

	greetings "github.com/hlozadaccs/go-learning-greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
