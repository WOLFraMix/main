package main

import "fmt"

// iota - это ключевое слово, которое позволяет генерировать последовательные числа.
// Define the first set of color constants starting from iota + 1.
const (
	Red   = iota + 1 // Red = 1
	Green            // Green = 2
	Blue             // Blue = 3
)

// Define the second set of color constants starting from iota.
const (
	Yellow = iota // Yellow = 0
	Orange        // Orange = 1
	Violet        // Violet = 2
)

const (
	A = iota * 2
	B
	C
	D
	_
	F
)

// Main function is the program's entry point.
func main() {
	// Print the first set of color constants.
	fmt.Println(Red, Green, Blue)
	// Print the second set of color constants.
	fmt.Println(Yellow, Orange, Violet)
	fmt.Println(A, B, C, D, F)
}
