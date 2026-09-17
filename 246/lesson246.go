package main

import (
	"fmt"
	"math"
	"sort"
)

// Point — структура для хранения координат точки
type Point struct {
	x, y int
}

// Расстояние всегда должно быть положительным числом
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	// Слайс для хранения исходных точек городов
	cities := make([]Point, n)
	// Отдельные слайсы только для X и Y координат
	xs := make([]int, n)
	ys := make([]int, n)

	// Используем мапу для проверки, занята ли точка городом
	occupied := make(map[Point]bool)

	// Чтение входных данных: координаты всех городов
	for i := 0; i < n; i++ {
		fmt.Scan(&cities[i].x, &cities[i].y)
		xs[i] = cities[i].x
		ys[i] = cities[i].y
		occupied[cities[i]] = true // Помечаем точку как занятую
	}

	// Сортируем координаты для поиска медианы
	sort.Ints(xs)
	sort.Ints(ys)

	// Находим медиану
	// Медиана обеспечивает минимальную сумму расстояний
	bestX := xs[n/2]
	bestY := ys[n/2]

	// Если оптимальная медианная точка свободна,
	// то это наш ответ
	if !occupied[Point{bestX, bestY}] {
		fmt.Println(bestX, bestY)
		return
	}

	// Если в медианной точке стоит город, ищем ближайшую свободную
	minSum := math.MaxInt64 // Минимальная найденная сумма расстояний
	resX, resY := 0, 0      // Координаты лучшего свободного места

	// Радиус поиска берем с запасом.
	// Обычно свободное место находится очень близко к медиане.
	radius := 10

	// Обход квадрата вокруг медианной точки
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			cx := bestX + dx
			cy := bestY + dy

			p := Point{cx, cy}

			// Пропускаем точки, которые уже заняты городами
			if occupied[p] {
				continue
			}

			// Считаем суммарное расстояние до всех городов
			currentSum := 0
			for _, city := range cities {
				currentSum += abs(cx-city.x) + abs(cy-city.y)
			}

			// Если нашли точку с меньшей суммой расстояний, запоминаем её
			if currentSum < minSum {
				minSum = currentSum
				resX = cx
				resY = cy
			}
		}
	}

	// Выводим координаты наилучшей свободной точки
	fmt.Println(resX, resY)
}
