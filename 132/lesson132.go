package main

import "fmt"

func twoSum(arr []int, sum int) []int {
	left := 0
	right := len(arr) - 1

	for left != right {
		tmp := arr[left] + arr[right]

		if tmp == sum {
			return []int{arr[left], arr[right]}
		}

		// движение указателей
		if tmp < sum {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main() {
	arr := []int{1, 2, 3, 4, 6}
	target := 8
	result := twoSum(arr, target)

	if result != nil {
		fmt.Printf("Пара найдена: %d, %d\n", result, result)
	} else {
		fmt.Println("Пара не найдена")
	}
}
