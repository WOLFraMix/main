package main

import (
	"fmt"
	"log"
)

// Initializes two float64 values, calls calculate, and prints the result or logs a fatal error.
func main() {
	a, b := 10.0, 0.0
	result, err := calculate(a, b)
	if err != nil {
		log.Fatalf("Unable to calculate: %s", err)
	} else {
		fmt.Println("result is:", result)
	}
}

// calculate executes a sequence of operations on two float64 numbers.
// It first runs logicX, then performs division via divide, returning the result or an error.
func calculate(a, b float64) (float64, error) {
	if err := logicX(); err != nil {
		return 0, fmt.Errorf("LogicX: %w", err) // %w - для ошибок
	}

	result, err := divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("Divide: %w", err)
	}

	return result, nil
}

// logicX represents a custom logic step.
// Currently, it always succeeds and returns nil.
func logicX() error {
	return nil
}

// divide performs division of a by b.
// It checks for division by zero and returns an error if b is 0.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}
