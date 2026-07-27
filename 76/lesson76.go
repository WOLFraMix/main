package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
	printTable(5)
}

func printTable(num int) {
	for x := 1; x <= num; x++ {
		for y := 1; y <= num; y++ {
			fmt.Printf("%d x %d = %d\t", x, y, x*y)
		}
		fmt.Println()
	}
}
