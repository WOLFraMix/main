package main

// функция-помощник getOSName
// Инкапсулирует логику определения имени ОС
// Делая код модульным
import (
	"fmt"
	"runtime"
)

func getOSName() string { // сначала создаём функцию
	switch runtime.GOOS {
	// switch - конструкция для выбора варианта в зависимости от значения
	// runtime.GOOS - предопределённая константа, содержащая имя текущей ОС
	case "darwin": // даём нормальное имя для MacOS
		return "MacOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default: // если ОС не распознана, возвращаем строку с дефолтным значением
		return fmt.Sprintf("Unknown OS (%s)", runtime.GOOS)
	}
}

func main() { // теперь в main() просто вызываем функцию
	fmt.Println("This test will analyze your operating system")
	fmt.Println("Your OS is:", getOSName())
}
