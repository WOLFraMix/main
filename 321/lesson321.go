package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	inputFile := "stepik_io_2.text"
	outputFile := "result.txt"

	src, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при открытии исходного файла: %v\n", err)
		return
	}
	defer src.Close()

	dst, err := os.Create(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при создании целевого файла: %v\n", err)
		return
	}
	defer dst.Close()

	buf := make([]byte, 4096)

	for {
		n, readErr := src.Read(buf)
		if readErr != nil {
			if readErr == io.EOF {
				break // достигли конца файла — всё ок
			}
			fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", readErr)
			return
		}

		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			r := rune(buf[i])
			if unicode.IsUpper(r) {
				continue // пропускаем заглавные буквы
			}
			_, writeErr := dst.Write([]byte{buf[i]})
			if writeErr != nil {
				fmt.Fprintf(os.Stderr, "Ошибка записи: %v\n", writeErr)
				return
			}
		}
	}

	fmt.Println("Готово: файл создан без заглавных букв —", outputFile)
}
