package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"
)

func main() {
	result, err := run()
	if err != nil {
		log.Fatalf("Ошибка:", err)
	}
	fmt.Println("Результат работы:", result)
}

func run() (val string, err error) {
	if x := rand.IntN(100); x < 50 {
		val, err = createNewValue(x)
		if err != nil {
			return "", fmt.Errorf("Создание значения с %d: %w", x, err)
		}
	} else {
		val, err = createDefaultValue()
		if err != nil {
			return "", fmt.Errorf("Создание дефолтного значения: %w", err)
		}
	}
	return val, nil
}

func createNewValue(x int) (string, error) {
	return strconv.Itoa(x), nil
}

func createDefaultValue() (string, error) {
	return "-1", nil
}
