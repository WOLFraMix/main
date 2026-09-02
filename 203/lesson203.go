package main

import "fmt"

// findConnectedComponents находит все компоненты связности в графе
func findConnectedComponents(graph map[int][]int) [][]int {
	visited := make(map[int]bool)
	var connectedComponents [][]int

	// Инициализируем все вершины как непосещённые
	for v := range graph {
		visited[v] = false
	}

	// Проходим по всем вершинам графа
	for v := range graph {
		if !visited[v] {
			var component []int
			dfs(graph, v, visited, &component)
			connectedComponents = append(connectedComponents, component)
		}
	}

	return connectedComponents
}

// dfs выполняет поиск в глубину и заполняет текущую компоненту связности
func dfs(graph map[int][]int, v int, visited map[int]bool, component *[]int) {
	visited[v] = true
	*component = append(*component, v)

	for _, neighbor := range graph[v] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited, component)
		}
	}
}

func main() {
	// Граф из примера на изображении
	graph := map[int][]int{
		1:  {2, 3},
		2:  {1, 3},
		3:  {1, 2},
		4:  {6, 7},
		5:  {6, 7},
		6:  {4, 5, 7},
		7:  {4, 5, 6},
		8:  {11},
		9:  {10, 11},
		10: {9},
		11: {8, 9},
	}

	components := findConnectedComponents(graph)

	fmt.Println("Компоненты связности:")
	for i, comp := range components {
		fmt.Printf("Компонента %d: %v\n", i+1, comp)
	}
}
