package main

// делаем вариант с заменой логики
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This test will analyze your operating system") // заголовок

	os := runtime.GOOS // сохраняем ОС в переменную
	if os == "darwin" {
		os = "MacOS"
	}
	fmt.Println("Your OS is:", os) // выводим результат
}
