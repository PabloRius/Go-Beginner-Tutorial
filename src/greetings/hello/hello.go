package main

import (
	"fmt"
	"log"

	"greetings.com/greetings"
)

func main() {
	// message := greetings.Hello("Pablo")
	// fmt.Println(message)

	// GoLang allows you to set some flags to the default logger
	log.SetPrefix("Greetings: ")
	log.SetFlags(0) // Won't log the time, source file or line number
	// message, err := greetings.Hello("") 		// The function would return an error
	message, err := greetings.Hello("Pablo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	names := []string{"Pablo", "Serena"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
