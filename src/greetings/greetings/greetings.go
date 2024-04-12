package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// Long variable declaration
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	// Shorter declaration, var tyep is inferred from the right side of the declaration
	// message := fmt.Sprintf("Hi, %v. Welocme!", name)
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomPhrase(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomPhrase() string {
	phrases := []string{
		"Hi, %v. Welcome!",
		"Great to see you , %v!",
		"Hola %v amigo!",
	}
	return phrases[rand.Intn(len(phrases))]
}
