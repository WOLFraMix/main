package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
)

func main() {

	result, err := calculate(6, 0)

	if err != nil {
		// ошибка в main
		log.Fatalf("calculate: %v", err)
	}

	fmt.Println("Результат работы", result)

}

func calculate(num1, num2 int) (int, error) {
	if rand.IntN(100) > 50 {
		return 0, errors.New("здесь возникла ошибка")
	}

	result, err := divide(num1, num2)

	if err != nil {
		// ошибка оборачивающая
		return 0, fmt.Errorf("divide: %w", err)
	}

	return result, nil

}

func divide(a, b int) (int, error) {

	if b == 0 {
		// ошибка итоговая
		return 0, errors.New("divide by zero")
	}

	return a / b, nil

}
