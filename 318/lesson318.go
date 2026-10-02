package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	m := make(map[string]int)
	parts := strings.Fields(input)

	for _, word := range parts {
		m[word]++
	}

	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}
