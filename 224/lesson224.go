package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите строку: ")
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatalf("Ошибка при чтении ввода: %s", err)
		} else {
			log.Fatalf("Не удалось считать строку.")
		}
	}
	fmt.Println("Вы ввели:", scanner.Text())
}
