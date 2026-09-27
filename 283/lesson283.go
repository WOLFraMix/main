package main

import "fmt"

type Logger interface {
	Log(msg string)
}

type EmailLogger struct{}

func (e *EmailLogger) Log(msg string) {
	fmt.Printf("Отправка на почту: %s\n", msg)
}

type FileLogger struct{}

func (f *FileLogger) Log(msg string) {
	fmt.Printf("Запись в файл: %s\n", msg)
}

func NewLogger(useEmail bool) Logger {
	if useEmail {
		return &EmailLogger{}
	}
	return nil
}

func main() {
	var logger1 Logger = NewLogger(true)
	if logger1 != nil {
		logger1.Log("Система запущена")
	} else {
		fmt.Println("Логгер не инициализирован, пропускаем запись")
	}

	var logger2 Logger = NewLogger(false)
	if logger2 != nil {
		logger2.Log("Система запущена")
	} else {
		fmt.Println("Логгер не инициализирован, пропускаем запись")
	}
}
