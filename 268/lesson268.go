package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		words := strings.Fields(line)

		for _, word := range words {
			if []rune(word)[0] == 'a' {
				fmt.Println(word)
			}
		}

	}
}
