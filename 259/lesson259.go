package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, _ := strconv.Atoi(line)

	data := make(map[string]int, n)

	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)

	for i := 0; i < n; i++ {
		key := parts[i]
		data[key] = 1
	}

	var sum int
	for _, v := range data {
		sum += v
	}

	fmt.Println(sum)
}
