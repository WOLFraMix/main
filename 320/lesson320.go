package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Читаем весь файл целиком
	data, err := os.ReadFile("stepik_io_1.text")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения файла: %v\n", err)
		return
	}

	sum := 0
	numStr := ""

	for _, ch := range string(data) {
		if ch >= '0' && ch <= '9' {
			// Если символ — цифра, добавляем к текущей последовательности
			numStr += string(ch)
		} else {
			// Если не цифра и у нас уже есть накопленное число — конвертируем и добавляем к сумме
			if numStr != "" {
				num, errConv := strconv.Atoi(numStr)
				if errConv == nil {
					sum += num
				}
				numStr = "" // сбрасываем накопленную строку
			}
		}
	}

	// Если строка заканчивается числом (нет завершающего не-цифрового символа)
	if numStr != "" {
		num, errConv := strconv.Atoi(numStr)
		if errConv == nil {
			sum += num
		}
	}

	fmt.Println(sum)
}
