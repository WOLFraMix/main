package main

import "fmt"

var input string

func main() {
	scanUserInput()
	sayHello()
}

func scanUserInput() {
	fmt.Println("Enter your name: ")
	fmt.Scanln(&input)
}

func sayHello() {
	fmt.Printf("Hello, %s!\n", input)
}
