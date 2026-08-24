package main

import (
	"cmp" // для функции cmp.Compare (сравнение значений)
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	movies := map[string]map[string]float64{
		"Экшен": {
			"Убить Билла":   8.52,
			"Отчаяние":      5.89,
			"Нейрослоп3000": 0.0001,
			"Интерстеллар":  9.1,
			"AmogUs":        10.00,
		},
		"Драма": {
			"Drama Queen":            7.524,
			"Поющие в терновнике":    7.527,
			"Титаник":                9.54,
			"Возвращение драмозавра": 9.51,
			"Касабланка":             7.0001,
			"Спячка":                 6.99999,
		},
		"Комедия": {
			"BadComedian":          6.9090909,
			"GoodComedian":         7.000000001,
			"Ржущие кони":          9.99,
			"Быстрая Сортировка":   7.5,
			"Сортировка слиянием":  7.5,
			"Сортировка пузырьком": 7.5,
		},
		"Ужасы": {
			"Месть Эль Диабло":               0.33,
			"Дикие орангутанги разносят код": 1.5,
		},
	}
	// Выводим рекомендации по фильмам (с фильтрацией и сортировкой)
	printRecommendations(movies)
}

// printRecommendations выводит рекомендации фильмов по жанрам и рейтингу
func printRecommendations(movies map[string]map[string]float64) {
	// Оставляем только фильмы с рейтингом >= 7.0
	filtered := filterByRating(movies, 7.0)
	// Получаем список жанров и сортируем их по алфавиту
	genres := slices.Sorted(maps.Keys(filtered))
	for _, genre := range genres {
		films := filtered[genre]
		// Сортируем названия фильмов: сначала по рейтингу, потом по имени
		sorted := sortedFilmNames(films)
		// Формируем строки вида «Название (рейтинг)»
		formatted := formatFilms(films, sorted)
		// Печатаем строку для жанра: «Жанр: фильм1 (рейтинг), ...»
		fmt.Printf("%s: %s.\n", genre, strings.Join(formatted, ", "))
	}
}

// filterByRating возвращает новую мапу с рейтингом не ниже minRating
func filterByRating(movies map[string]map[string]float64, minRating float64) map[string]map[string]float64 {
	result := make(map[string]map[string]float64)
	for genre, films := range movies {
		filteredFilms := make(map[string]float64)
		for film, rating := range films {
			if rating >= minRating {
				filteredFilms[film] = rating
			}
		}
		// Добавляем жанр в результат, только если в нём есть подходящие фильмы
		if len(filteredFilms) > 0 {
			result[genre] = filteredFilms
		}
	}
	return result
}

// sortedFilmNames возвращает отсортированный слайс фильмов
func sortedFilmNames(films map[string]float64) []string {
	return slices.SortedFunc(maps.Keys(films), func(a, b string) int {
		// Сравниваем рейтинги и названия
		return cmp.Or(
			cmp.Compare(films[b], films[a]),
			strings.Compare(a, b))
	})
}

// formatFilms формирует слайс строк вида «Название (5.0 рейтинг)»
func formatFilms(films map[string]float64, names []string) []string {
	result := make([]string, len(names))
	for i, name := range names {
		rating := films[name]
		result[i] = fmt.Sprintf("%s (%.1f)", name, rating)
	}
	return result
}
