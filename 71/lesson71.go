package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
	}
	// тоже самое что:
	for range 5 {
		fmt.Println("Привет")
	}

	rollDice(5)
}

func rollDice(x int) {
	var z, y, i int
	for z+y != x {
		z = rand.IntN(6) + 1
		y = rand.IntN(6) + 1
		i++

		if z+y != x {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n", z, y, z+y)
			continue
		}
		if z+y == x {
			var word1 string = plural1(i)
			var word2 string = plural2(i)
			fmt.Printf("Выпало %d и %d, в сумме %d, на это %s %d %s.\n", z, y, x, word1, i, word2)
			continue
		}
	}
}

func plural1(n int) string {
	lastLetter := n % 10
	lastLetters := n % 100

	if lastLetters == 11 {
		return "потребовалось"
	}
	if lastLetter == 1 {
		return "потребовался"
	}
	return "потребовалось"
}

func plural2(n int) string {
	lastLetter := n % 10
	lastLetters := n % 100

	if lastLetters >= 11 && lastLetters <= 19 {
		return "бросков"
	}
	if lastLetter == 1 {
		return "бросок"
	}
	if lastLetter == 2 || lastLetter == 3 || lastLetter == 4 {
		return "броска"
	}
	return "бросков"
}
