package main

import "fmt"

// Вариант 1: принимает срез (копия заголовка среза)
func modifySliceValue(s []int) {
	s[0] = 999        // ✅ изменит элемент (общий массив)
	s = append(s, 42) // ❌ НЕ повлияет на оригинальный срез!
}

// Вариант 2: принимает указатель на срез
func modifySlicePointer(s *[]int) {
	(*s)[0] = 888       // ✅ изменит элемент
	*s = append(*s, 42) // ✅ изменит оригинальный срез!
}

func main() {
	data := []int{1, 2, 3}
	fmt.Println("Исходный:", data) // [1 2 3]

	modifySliceValue(data)
	fmt.Println("После modifySliceValue:", data) // [999 2 3] — append не сработал!

	modifySlicePointer(&data)
	fmt.Println("После modifySlicePointer:", data) // [888 2 3 42] — append сработал!
}
