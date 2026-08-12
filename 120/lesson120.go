package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	k := 9
	result1 := twoSum(nums, k)
	fmt.Println(result1) // [0 1]
	// тоже самое что:
	result2 := find(nums, k)
	fmt.Println(result2)
}

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

func find(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// если не нашли пары подходящих чисел
	return nil
}
