package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for i := 0; i <= 100; i++ {
		scanner.Scan()
		if scanner.Text() == "" {
			break
		}
		if scanner.Err() != nil {
			break
		}
		input := scanner.Text()
		n, err := strconv.Atoi(input)
		if err != nil {
			break
		}
		sum += n
	}
	result := strconv.Itoa(sum)
	io.WriteString(os.Stdout, result)
}
