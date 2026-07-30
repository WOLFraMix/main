package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{5, 6, 7, 1, 1}
	sl3 := []int{2, 0, -3}
	sl4 := []int{}
	sl5 := []int{5}

	printMagic(sl1)
	printMagic(sl2)
	printMagic(sl3)
	printMagic(sl4)
	printMagic(sl5)
}

func printMagic(slice []int) {
	if len(slice) <= 0 {
		fmt.Println("[]")
		return
	}
	if len(slice) <= 1 {
		fmt.Println("[1]")
		return
	}

	result := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		num := 1
		for j := 0; j < len(slice); j++ {
			if i == j {
				continue
			}
			num *= slice[j]
		}
		result[i] = num
	}

	// Преобразуем числа в строки
	str := make([]string, len(result))
	for i, v := range result {
		str[i] = strconv.Itoa(v)
	}

	// Формируем строку вида [24, 12, 8, 6]
	output := "[" + strings.Join(str, ", ") + "]"
	fmt.Println(output)
}
