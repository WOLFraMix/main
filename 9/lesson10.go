package main

import "fmt"

func main() {
	myArray := []int{7, 8, 9} // создаём массив из 3х элементов

	for i := 0; i < len(myArray); i++ { // проходим по каждому элементу массива
		item := myArray[i] // сохраняем элемент в переменную item
		fmt.Println(item)  // выводим значение переменной item на экран
	}
}
