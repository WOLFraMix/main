package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println(nums)

	var numbers = [5]int{2, 4, 6, 8, 10}
	fmt.Println(numbers)

	fruits := [...]string{"apple", "banana", "cherry", "orange"}
	fmt.Println(fruits)

	numb := [10]int{0: 5, 1: 10, 4: 20, 7: 30} // отсчёт в массиве идёт с 0 индекса
	fmt.Println(numb)

	// var nums [5]int
	// var nums = [5]int{}
	// nums := [5]int{}
}
