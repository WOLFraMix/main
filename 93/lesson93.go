package main

import "fmt"

func main() {
	t1 := []int{1, 2, 3}
	t2 := []int{4, 5, 6}
	t3 := []int{10, 20, 30, 40}
	t4 := []int{1, 2}
	t5 := []int{}
	t6 := []int{1, 2, 3, 4}
	t7 := []int{}
	t8 := []int{}
	t9 := []int{-1, 2, -3}
	t0 := []int{4, -5, 6}

	fmt.Println(SumSlices(t1, t2))
	fmt.Println(SumSlices(t3, t4))
	fmt.Println(SumSlices(t5, t6))
	fmt.Println(SumSlices(t7, t8))
	fmt.Println(SumSlices(t9, t0))
}

func SumSlices(s1 []int, s2 []int) (result []int) {
	// сначала проверяем на длину слайсы и возвращаем пустой слайс
	result = []int{}
	if len(s1) == 0 || len(s2) == 0 {
		return result
	}

	// делаем длину результата минимальной
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}
	result = make([]int, minLen)

	// делаем цикл и вносим значения в результат
	for i := 0; i < minLen; i++ {
		result[i] = s1[i] + s2[i]
	}
	return result
}
