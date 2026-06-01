package main

import "fmt"

func sayHello() {
	fmt.Println("Hello!")
}

func greet(name string) {
	fmt.Println("Hello,", name+"!")
}

func add(a, b int) int {
	return a + b
}

func main() {
	sayHello()
	greet("Alice")
	greet("Bob")
	result := add(5, 10)
	fmt.Println("5 + 10 =", result)
}
