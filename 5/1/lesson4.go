package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This test will analyze your operating system")
	fmt.Println("Your OS is:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	case "freebsd":
		fmt.Println("FreeBSD")
	default:
		fmt.Printf("%s\n", os)
	}
}
