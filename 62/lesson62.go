package main

import "fmt"

func main() {
	fn(5)
}

func fn(count int) {
	if count <= 0 {
		return
	}
	fmt.Println(count)
	fn(count - 1)
}
