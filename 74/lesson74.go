package main

import "fmt"

func main() {
	count := 0

	for {
		fmt.Println("Счётчик:", count)
		count++
		if count >= 5 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("Значение i:", i)
	}
}

// break - завершает цикл
// continue - прекращает только текущую итерацию
// return - возвращает значение и завершает всю функцию
