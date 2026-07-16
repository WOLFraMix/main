package main

import "fmt"

func main() {
	fmt.Println(checkValue(5))
}

func checkValue(x int) string {
	if x < 0 {
		return "Отрицательное число"
	}
	if x == 0 {
		return "Ноль"
	}
	if x <= 10 {
		return "Небольшое положительное число"
	}
	return "Большое положительное число"

}
