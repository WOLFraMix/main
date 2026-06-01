package main

// оптимизированный вариант
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This test will analyze your operating system")
	fmt.Println("Your OS is:", runtime.GOOS)
}
