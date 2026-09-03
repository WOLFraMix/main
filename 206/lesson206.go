package main

import (
	"errors"
	"fmt"
)

func main() {
	err := someFunction()
	if err != nil {
		return
	}
	err = fn()
	if err != nil {
		return
	}
}

func someFunction() error {
	return errors.New("это сообщение с ошибкой")
}

func fn() error {
	return fmt.Errorf("ошибка в функции %s", "fn")
}
