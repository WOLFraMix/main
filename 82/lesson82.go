package main

import "fmt"

func main() {
	nums := [5]int{5, 23, -9, 1, 6}

	for i := 0; i <= len(nums)-1; i++ {
		fmt.Println(i, nums[i])
	}
	// тоже самое что:
	for index, value := range nums {
		fmt.Println(index, value)
	}

	test1 := [10]int{4, 2, 3, 3, 9, 9, 3, 3, 2, 4}
	test2 := [10]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 1}

	isPalindrome(test1)
	isPalindrome(test2)
}

func isPalindrome(arr [10]int) { // палиндром = зеркальный
	for i := 0; i < len(arr)/2; i++ { // кол-во итераций = половина длины массива
		if arr[i] != arr[len(arr)-i-1] { // сначала проверяем НЕ совпадения
			fmt.Println("He палиндром!") // тогда при первом же НЕ совпадении
			return                       // функция вернёт ответ сразу
		}
	}
	fmt.Println("Это палиндром!") // иначе другой ответ
}
