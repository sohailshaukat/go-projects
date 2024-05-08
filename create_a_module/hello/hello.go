package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samanth", "Darrin"}
	// fmt.Println(quote.Go())
	// message, err := greetings.Hello("Gladys")
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

// to track dependencies -> go mod init example.com/<project-name>
// to run -> go run .
// to install dependencies -> go mod tidy
