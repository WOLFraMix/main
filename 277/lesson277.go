package main

import "fmt"

// Функция, принимающая другую функцию как callback
func doMath(a, b int, callback func(int)) {
	result := a + b
	callback(result) // вызываем callback с результатом
}

func main() {
	// Вызываем doMath с анонимной функцией, которая выводит результат
	doMath(2, 3, func(res int) {
		fmt.Printf("Результат: %d\n", res)
	})
}
