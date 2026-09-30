package main

import (
	"fmt"
	"io"
	"strings"
)

func sumDigits(n int) int {
	if n < 0 {
		n = -n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(sumDigits(10))

	r := strings.NewReader("Привет")
	processData(r)

	a := []int{5, 5, 1}
	b := []int{1, 3, 5}
	fmt.Println(union(a, b)) // Вывод: [5 1 3]
}

func processData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		n, err := reader.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %v bytes: %v\n", n, string(b[:n]))
	}
}

// union объединяет два слайса int, оставляя только уникальные элементы.
// Порядок: сначала все уникальные элементы из first (в исходном порядке),
// затем уникальные элементы из second, которых не было в first.
func union(a, b []int) []int {
	seen := make(map[int]struct{})
	var result []int

	// Добавляем элементы из первого слайса
	for _, v := range a {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	// Добавляем элементы из второго слайса, если их ещё не было
	for _, v := range b {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
