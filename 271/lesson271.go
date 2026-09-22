package main

import "fmt"

func main() {
	var s string
	var max string

	for {
		fmt.Scan(&s)
		if s == "stop" {
			fmt.Println("Самое длинное слово:", max)
			return
		}
		if len(s) > len(max) {
			max = s
		}
	}
}
