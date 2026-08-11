package main

import "fmt"

// изображение — это массив из 1080 строк длиной в 1920 пикселей
// каждый пиксель — массив из трёх байт
// 1080 — размер массива
// [1920][3]uint8 — тип элемента
var rgbImage [1080][1920][3]uint8

func main() {
	// 3-я строка в изображении
	line := rgbImage[2]

	// 4-й пиксель в третьей строке изображения
	pixel := rgbImage[2][3]

	// значение синей компоненты (второй байт) 4-го пикселя в третьей строке изображения
	red := rgbImage[2][3][1]

	fmt.Println(cap(rgbImage), cap(line), cap(pixel), red)

	// средняя температура ежедневно в неделе
	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

	// сумма средних температур за неделю
	sumTemp := 0

	for i := 0; i < len(weekTemp); i++ {
		sumTemp += weekTemp[i]
	}

	// средняя температура за неделю
	average := sumTemp / len(weekTemp)
	fmt.Println(average)
}
