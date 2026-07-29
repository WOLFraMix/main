package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [5]int{3, 8, 1, 8, 1}
	fmt.Println(secretGen(arr))
}

func secretGen(array [5]int) (result string) {
	// генерируем "секретный" ключ
	s1 := array[0] // находим минимальное значение
	for i := range array {
		if array[i] < s1 {
			s1 = array[i]
		}
	}
	s3 := array[0] // находим максимальное значение
	for i := range array {
		if array[i] > s3 {
			s3 = array[i]
		}
	}
	s2 := "" // вставляем буквы перед числами
	for i := range array {
		switch array[i] % 2 {
		case 0: // "E" (от слова Even — четный)
			s2 = s2 + "E" + strconv.Itoa(array[i])
		default: // "O" (от слова Odd — нечетный)
			s2 = s2 + "O" + strconv.Itoa(array[i])
		}
	}
	// собираем строку
	result = strconv.Itoa(s1) + s2 + strconv.Itoa(s3)
	return
}
