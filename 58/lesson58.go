package main

import (
	"fmt"
	"log"
)

// Demonstrates the usage of the helloFactory function by creating
// greeting functions for Russian, English, and French, and then calling them.
func main() {
	// Create a Russian greeting function and handle any potential errors.
	ruFn, err := helloFactory("ru")
	if err != nil {
		log.Fatalf("helloFactory error: %s", err.Error())
	}
	// Call the Russian greeting function with the name "Катя".
	ruFn("Катя")

	// Create an English greeting function and handle any potential errors.
	enFn, err := helloFactory("en")
	if err != nil {
		log.Fatalf("helloFactory error: %s", err.Error())
	}
	// Call the English greeting function with the name "Stepan".
	enFn("Stepan")

	// Create a French greeting function and handle any potential errors.
	frFn, err := helloFactory("fr")
	if err != nil {
		log.Fatalf("helloFactory error: %s", err.Error())
	}
	// Call the French greeting function with the name "Emma".
	frFn("Emma")
}

// helloFactory returns a closure that prints a greeting in the specified language.
// It takes a language code ("en", "ru", etc.) and returns a function that accepts a name,
// along with an error if the language is not supported.
func helloFactory(lang string) (func(name string), error) {
	var message string
	// Determine the greeting format based on the provided language.
	switch lang {
	case "en":
		message = "Hello, %s!\n"
	case "ru":
		message = "Привет, %s!\n"
	default:
		// Return an error if the language is not recognized.
		return nil, fmt.Errorf("unknown language: %s", lang)
	}
	// Return a closure that formats and prints the greeting using the selected message.
	return func(name string) {
		fmt.Printf(message, name)
	}, nil
}
