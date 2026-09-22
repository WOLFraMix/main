package main

import "fmt"

func main() {
	var capacity int
	fmt.Scan(&capacity)
	queue := make([]int, 0, capacity)

	for i := 0; i < 10; i++ {
		if len(queue) == capacity {
			queue = queue[1:]
		}
		queue = append(queue, i)
		fmt.Printf("Добавлен %d: ", i)
		fmt.Print(queue)
		fmt.Println()
	}
}
