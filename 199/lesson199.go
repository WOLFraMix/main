package main

import (
	"fmt"
)

// Функция BFS для графа, представленного как map[вершина][]соседей
func bfs(graph map[int][]int, start int) []int {
	// visited хранит посещённые вершины
	var visited = make(map[int]bool)
	// queue — очередь вершин для обхода
	queue := []int{start}
	// result будет содержать порядок посещения вершин
	result := []int{}

	for len(queue) > 0 {
		// Извлекаем первую вершину из очереди
		currentVertex := queue[0]
		// Удаляем её из начала очереди
		queue = queue[1:]

		// Если вершина ещё не была посещена
		if !visited[currentVertex] {
			// Добавляем её в результат и отмечаем как посещённую
			result = append(result, currentVertex)
			visited[currentVertex] = true

			// Обрабатываем соседей текущей вершины
			for _, neighbor := range graph[currentVertex] {
				// Если сосед не был посещён ранее
				if !visited[neighbor] {
					// Добавляем его в конец очереди
					queue = append(queue, neighbor)
				}
			}
		}
	}
	return result
}

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {},
	}

	startingNode := 1
	order := bfs(graph, startingNode)
	fmt.Println("Порядок обхода:", order)
}
