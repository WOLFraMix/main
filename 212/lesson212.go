package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	// Чтение ввода:
	// n k
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}

	shuffleKSlices(slice, k)
	result := sliceToString(slice)
	fmt.Println(result)
}

// shuffleKSlices перемешивает слайс k раз.
func shuffleKSlices(slice []int, k int) {
	l := len(slice)
	for i := 0; i < k; i++ {
		for j := l - 1; j > 0; j-- {
			p := rand.IntN(j + 1)
			slice[j], slice[p] = slice[p], slice[j]
		}
	}
}

// sliceToString превращает слайс чисел в строку.
func sliceToString(slice []int) string {
	parts := make([]string, len(slice))
	for i, v := range slice {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, " ")
}
