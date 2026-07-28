package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	modifyArray(nums)
	fmt.Println(nums) // оригинал массива без изменений

	modifyArray2(&nums)
	fmt.Println(nums) // меняем оригинал
}

func modifyArray(arr [5]int) { // сюда передаётся копия массива
	arr[2] = 100
	fmt.Println(arr)
}

func modifyArray2(arr *[5]int) { // сюда передаётся значение в памяти
	arr[2] = 200
	fmt.Println(arr)
}
