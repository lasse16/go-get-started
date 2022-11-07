package main

import (
	"fmt"
	"get-started/sender"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := sender.GreetMultiple(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
