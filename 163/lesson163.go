package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	m := map[string]map[string]float64{
		"Экшен": {
			"Фильм1": 8.52,
			"Фильм2": 6.0,
		},
		"Драма": {
			"Фильм3": 7.524,
			"Фильм4": 7.527,
			"Фильм5": 5.54,
		},
		"Трэш": {
			"Фильм11": 4.52,
			"Фильм22": 1.0,
		},
	}
	printRecommendations(m)
}

func printRecommendations(movies map[string]map[string]float64) {
	recommendations := make(map[string][]string)

	for genre, films := range movies {
		for film, rating := range films {
			if rating >= 7 {
				recommendations[genre] = append(recommendations[genre], film)
			}
		}
	}

	sortedGenres := []string{}
	for genre := range recommendations {
		sortedGenres = append(sortedGenres, genre)
	}
	slices.Sort(sortedGenres)

	for _, genre := range sortedGenres {
		slices.SortFunc(recommendations[genre], func(a, b string) int {
			ratingA := movies[genre][a]
			ratingB := movies[genre][b]

			if ratingA == ratingB {
				if a > b {
					return 1
				}
				if a < b {
					return -1
				}
				return 0
			}
			if ratingA > ratingB {
				return -1
			}
			if ratingA < ratingB {
				return 1
			}
			return 0
		})
		recommended := []string{}
		for _, film := range recommendations[genre] {
			recommended = append(recommended, fmt.Sprintf("%s (%.1f)", film, movies[genre][film]))
		}
		fmt.Printf("%s: %s.\n", genre, strings.Join(recommended, ", "))
	}

}
