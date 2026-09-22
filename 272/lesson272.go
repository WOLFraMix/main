package main

import "fmt"

func main() {
	m := make(map[string]int)
	var str string

	for {
		fmt.Scan(&str)
		if str == "stop" {
			fmt.Println(m)
			return
		}
		m[str] += 1
	}
}
