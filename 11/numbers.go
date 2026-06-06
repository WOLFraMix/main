package main

import "fmt"

func main() {
	number := 5
	secondNumber := 7
	thirdNumber := 10

	if number > secondNumber {
		fmt.Println("первое число больше второго")
	} else if number < secondNumber {
		fmt.Println("второе число больше первого")
	} else {
		fmt.Println("первое и второе число равны")
	}

	switch 15 {
	case number + secondNumber:
		fmt.Println("первый скилл и второй!")
	case number + thirdNumber:
		fmt.Println("первый скилл и третий!")
	}
}
