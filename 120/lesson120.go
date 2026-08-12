package main

import "fmt"

func twoSum(nums []int, k int) []int {
	seen := make(map[int]int) // value -> index

	for i, num := range nums {
		complement := k - num
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}

	return nil // пустой слайс, если пары нет
}

func main() {
	nums := []int{2, 7, 11, 15}
	k := 9
	result := twoSum(nums, k)
	fmt.Println(result) // [0 1]
}
