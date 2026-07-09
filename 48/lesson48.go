package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Calculate and print the sum of two integers
	result := sum(2, 5)
	fmt.Println(result)

	// Define a string and compute its length in bytes and runes
	helloStr := "Привет, друг!"
	bytes, runes := getFullLength(helloStr)
	fmt.Printf("Строка %q имеет длину %d байт и %d рун\n", helloStr, bytes, runes)
	// %q - это строка в кавычках

	// Generate and print a random compliment for the given name
	fmt.Println(generateCompliment("Катя"))
}

// sum returns the sum of two integers.
func sum(n1, n2 int) int {
	return n1 + n2
}

// getFullLength returns the length of a string in bytes and the number of runes (characters).
func getFullLength(str string) (int, int) {
	return len(str), len([]rune(str)) // длина строки и длина строки в рунах
}

// generateCompliment returns a randomly selected compliment string for the specified name.
func generateCompliment(name string) string {
	// Generate a random integer between 0 and 2
	num := rand.Intn(3)
	switch num {
	case 0:
		return fmt.Sprintf("Ты вызываешь восторг, %s!", name)
	case 1:
		return fmt.Sprintf("У тебя потрясающая улыбка, %s!", name)
	default:
		return fmt.Sprintf("Ты вдохновляешь, %s!", name)
	}
}
