package main

import (
	"fmt"
	"log"
)

type MyError struct {
	Code    int
	Message string
}

func main() {
	if err := someFunc(); err != nil {
		log.Fatalf("Ошибка: %s", err)
	}
}

func (e MyError) Error() string {
	return fmt.Sprintf("Ошибка %d: %s.\n", e.Code, e.Message)
}

func someFunc() error {
	return &MyError{
		Code:    500,
		Message: "внутренняя ошибка",
	}
}
