package main

import "fmt"

var num = 20

func main() {
	x := 10
	fmt.Println("x в начале:", x)
	fmt.Println("num в блоке main:", num)
	if x > 5 {
		y := 20 // y есть только в блоке if
		fmt.Println("y в блоке:", y, "x в блоке:", x)
	}

	{
		x := 30 // этот x есть только в этом блоке
		fmt.Println("x во втором блоке:", x)
	}

	fmt.Println("x в конце:", x)
	fn()
	num += 5
	fn()
}

func fn() {
	fmt.Println(num)
}
