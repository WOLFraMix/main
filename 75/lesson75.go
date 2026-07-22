// Package main implements a binary search guessing game with verification.
package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
)

// main runs 100 verification iterations to ensure the binary search algorithm
// correctly guesses a randomly generated number within 6 attempts.
func main() {
	for range 100 {
		random = rand.IntN(100) + 1 // generate target number in range [1, 100]
		guesses = 0                 // reset attempt counter for this iteration
		result := play()            // execute binary search to guess the number
		if result != random {
			fmt.Printf("Неверный ответ. Было загадано число %d, а в ответе получили число %d", random, result)
			os.Exit(-1)
		}
	}
}

// guesses tracks the number of attempts made in the current game.
var guesses int

// random stores the target number that the algorithm needs to guess.
var random int

// guess evaluates a proposed number against the target.
// Returns -1 if the target is smaller, 1 if larger, 0 if correct.
// Returns an error if the maximum allowed attempts (6) is exceeded.
func guess(num int) (int, error) {
	if guesses >= 6 {
		return 0, errors.New("too many attempts")
	}
	guesses++
	if num > random {
		return -1, nil
	}
	if num < random {
		return 1, nil
	}
	return 0, nil
}

// play implements a binary search algorithm to find the random number.
// It repeatedly narrows the search range based on feedback from guess().
func play() int {
	low, high := 1, 100 // диапазон [1, 100]

	for {
		x := (low + high) / 2
		val, err := guess(x)
		if err != nil {
			return x
		}

		if val == 0 {
			return x
		}
		if val == -1 {
			// загаданное < x → сужаем верхнюю границу
			high = x - 1
		} else if val == 1 {
			// загаданное > x → сужаем нижнюю границу
			low = x + 1
		}
	}
}
