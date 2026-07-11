package main

import "fmt"

// Demonstrates calling the calculate function with different arithmetic operations.
func main() {
	// Perform addition
	result, err := calculate(1, 2, "add")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Perform subtraction
	result, err = calculate(1, 2, "subtract")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Perform multiplication
	result, err = calculate(1, 2, "multiply")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Perform division (note: b is 0, so it will trigger an error)
	result, err = calculate(10, 0, "divide")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// calculate performs a basic arithmetic operation on two float64 numbers.
// Parameters:
//   a: the first operand
//   b: the second operand
//   division: the operation to perform ("add", "subtract", "multiply", or "divide")
// Returns:
//   float64: the result of the operation
//   error: nil on success, or an error message for invalid operations or division by zero
func calculate(a float64, b float64, division string) (float64, error) {
	if division == "add" {
		return a + b, nil
	} else if division == "subtract" {
		return a - b, nil
	} else if division == "multiply" {
		return a * b, nil
	} else if division == "divide" && b == 0 {
		return 0, fmt.Errorf("division by zero")
	} else if division == "divide" {
		return a / b, nil
	} else {
		return 0, fmt.Errorf("unknown operation")
	}
}
