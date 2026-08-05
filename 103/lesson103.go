package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{3, 4, 5, 6, 7}
	s3 := []int{1, 3, 5, 7}
	s4 := []int{2, 4, 6, 8}
	var s5 []int
	s6 := []int{1, 2, 3}
	s7 := []int{}
	s8 := []int{}
	s9 := []int{-3, -2, -1, 0, 1, 2, 3}
	s0 := []int{-2, 0, 2, 4, 6}

	fmt.Println(intersectSlices(s1, s2))
	fmt.Println(intersectSlices(s3, s4))
	fmt.Println(intersectSlices(s5, s6))
	fmt.Println(intersectSlices(s7, s8))
	fmt.Println(intersectSlices(s9, s0))
}

func intersectSlices(slice1 []int, slice2 []int) (result []int, err error) {
	if slice1 == nil || slice2 == nil {
		return nil, fmt.Errorf("slices cannot be nil")
	}

	// создаём счётчики
	i, j := 0, 0

	for i < len(slice1) && j < len(slice2) {
		// ищем пересечения на отсортированных слайсах
		if slice1[i] == slice2[j] {
			// одинаковые значения добавляем в результат
			result = append(result, slice1[i])
			i++
			j++
		} else if slice1[i] < slice2[j] {
			i++
		} else {
			j++
		}
	}

	return result, nil
}
