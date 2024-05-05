package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined logger, including the log entry prefix and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())
	message, err1 := greetings.Hello("Harsh")

	// if an error was returned, print it to the console and exit the process
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(message)

	// A slice of names
	names := []string{"Abc", "Def", "Xyz"}

	messages, err2 := greetings.Hellos(names)

	if err2 != nil {
		log.Fatal(err2)
	}

	// If no errors was returned, print the returned message to the console
	fmt.Println(messages)
}