package main

import "fmt"

func main() { //   0  1  2  3  4
	nums := [5]int{2, 4, 6, 8, 0}
	fmt.Println(nums[1])

	nums[4] = 1000
	fmt.Println(nums)

	fmt.Println(len(nums))

	i := -1
	if i >= 0 && i <= len(nums)-1 {
		fmt.Println(nums[i])
	} else {
		fmt.Printf("Индекс %d выходит за пределы массива\n", i)
	}

	i = 3
	if i >= 0 && i <= len(nums)-1 {
		fmt.Println(nums[i])
	} else {
		fmt.Printf("Индекс %d выходит за пределы массива\n", i)
	}
}
