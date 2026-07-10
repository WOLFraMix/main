package main

import "fmt"

func main() {
	defer handlePanic() // Register a deferred function to recover from panics
	riskFunc()          // Call a function that intentionally panics
	fmt.Println("Эта строка не будет достигнута если произойдёт паника!")
}

func handlePanic() {
	if err := recover(); err != nil { // Recover from any panic and capture the error
		fmt.Println("Произошла паника:", err) // Print the panic message if a panic occurred
	}
}

func riskFunc() {
	panic("Что-то пошло не так") // Intentionally trigger a panic with a custom message
}
