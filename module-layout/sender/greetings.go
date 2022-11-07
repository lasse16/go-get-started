package sender

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name given")
	}
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func randomGreeting() string {
	greetings := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return greetings[rand.Intn(len(greetings))]

}

func init() {
	rand.Seed(time.Now().UnixNano())
}
