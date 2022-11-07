package main

import (
	"fmt"
	"get-started/sender"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := sender.Hello("TEST")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
