package main

import (
	"fmt"
	"log"

	greetings "github.com/hlozadaccs/go-learning-greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
