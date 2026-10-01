package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	runes := []rune(input)
	if len(runes) == 0 {
		fmt.Println("")
		return
	}

	var result []byte
	count := 1

	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++
		} else {
			result = append(result, fmt.Sprintf("%c%d", runes[i-1], count)...)
			count = 1
		}
	}
	result = append(result, fmt.Sprintf("%c%d", runes[len(runes)-1], count)...)

	fmt.Println(string(result))
}
