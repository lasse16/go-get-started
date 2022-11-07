package sender

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name given")
	}
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func GreetMultiple(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Greet(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
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
