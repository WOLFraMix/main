package main

import "fmt"

func main() {
	myName := "Slim Shady"
	fmt.Println("Hi, my name is... What? My name is... Who? My name is... chika chika... ",
		myName)

	name := "Stepan"
	fmt.Println("name = ", name)

	userName := "Admin"
	fmt.Println("username = ", userName)

	var sum int
	fmt.Println("the sum is = ", sum)

	a, b := 5, 10
	fmt.Println("a = ", a, "b = ", b)

	sum = a + b
	fmt.Println("sum = a + b - here is your answer for sum reason xD = ", sum)

	sum = sum + a + b
	fmt.Println("sum = sum + a + b - and now the sum is = ", sum)
}
