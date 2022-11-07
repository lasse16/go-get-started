package main

import (
	"fmt"
	"get-started/sender"
)

func main() {
	message := sender.Hello("TEST")
	fmt.Println(message)
}
