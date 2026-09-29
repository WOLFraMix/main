package main

import "fmt"

func main() {
	// Отложенная анонимная функция для обработки паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Это паника!")
		} else {
			fmt.Println("Программа отработала без паники!")
		}
	}()
}
