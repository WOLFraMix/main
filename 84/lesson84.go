package main

import (
	"fmt"
)

func main() {
	test1 := [3]string{"яблоко", "банан", "вишня"}
	test2 := [3]string{"Голанг", "Программирование", "Язык"}
	test3 := [3]string{"АЕИОУ", "кхм", ""}

	CountVowelsInArray(test1)
	CountVowelsInArray(test2)
	CountVowelsInArray(test3)
}

func CountVowelsInArray(arr [3]string) {
	vowels := "аеёиоуыэюяАЕЁИОУЫЭЮЯ"
	count0 := 0
	count1 := 0
	count2 := 0

	for i, value := range arr {
		for _, char := range value {
			for _, v := range vowels {
				if i == 0 && char == v {
					count0++
				}
				if i == 1 && char == v {
					count1++
				}
				if i == 2 && char == v {
					count2++
				}
			}
		}

	}
	fmt.Println(count0, count1, count2)
}
