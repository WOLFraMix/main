package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	arrM := [3][2]int{
		{2, 1}, // 0
		{4, 3}, // 1
		{6, 5}, // 2
	}
	fmt.Println(arrM)

	fmt.Println(arrM[1][1])

	arrM[2] = [2]int{9, 9}
	fmt.Println(arrM)

	arrM[1][0] = -5
	fmt.Println(arrM)

	s := [][]int{
		{2, 4, 5, 9},
		{6, 1, 2},
		{8, -7, 4, 3, 0},
	}
	fmt.Println(s)

	s[1] = append(s[1], 999)
	fmt.Println(s)

	s = append(s, []int{-5, -10})
	fmt.Println(s)

	sum := 0
	for _, innerS := range s {
		for _, v := range innerS {
			if v%2 != 0 && v < 0 { // В Go остаток от деления отрицательного числа может быть отрицательным.
				sum += v
			}
		}
	}
	fmt.Println(sum)
}
