package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Возводить в степень можно гораздо быстрее, чем за n умножений!
Реализуйте алгоритм быстрого возведения в степень.
*/

func readInput() (float64, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	return a, n
}

func exponentiation(a float64, n int) float64 {
	result := 1.0

	for n > 0 {
		if n%2 == 1 {
			result *= a
		}
		a *= a
		n /= 2
	}

	return result
}

func main() {
	a, n := readInput()

	result := exponentiation(a, n)
	fmt.Println(result)
}
