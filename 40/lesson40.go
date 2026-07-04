package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	num := rand.IntN(100)
	if num > 50 {
		fmt.Printf("Выпало число %d!\n", num)
	} else {
		fmt.Printf("Выпало число %d, маловато будет...\n", num)
	}
}
