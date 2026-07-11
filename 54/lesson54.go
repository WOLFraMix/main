package main

import (
	"fmt"
	"log"
)

// Tests the divide function with both valid and invalid inputs.
func main() {
	// First test: valid division
	result, err := divide(10, 5)
	if err != nil {
		log.Fatalf("error: %s", err)
	} else {
		fmt.Println(result)
	}
	// Second test: division by zero
	result, err = divide(10, 0)
	if err != nil {
		log.Fatalf("error: %s", err)
	} else {
		fmt.Println(result)
	}
}

// divide performs integer division of n1 by n2.
// Returns an error if n2 is zero.
func divide(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("деление на ноль")
	} else {
		return n1 / n2, nil
	}
}
