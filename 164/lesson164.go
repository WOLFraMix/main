package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	friendsData := map[string][]string{
		"Алексей":  {"Иван", "Сергей", "Елена"},
		"Иван":     {"Алексей", "Дмитрий", "Мария"},
		"Сергей":   {"Алексей", "Елена"},
		"Дмитрий":  {"Иван", "Елена", "Ольга"},
		"Елена":    {"Алексей", "Сергей", "Дмитрий"},
		"Мария":    {"Иван", "Ольга"},
		"Ольга":    {"Дмитрий", "Мария"},
		"Анна":     {"Петр"},
		"Петр":     {"Анна", "Сергей"},
		"Светлана": {"Иван", "Елена"},
	}

	fmt.Println(outputCountFriends(countFriends(friendsData)))
	fmt.Println(commonFriends(friendsData, "Алексей", "Дмитрий"))
	fmt.Println(mostPopularUsers(friendsData))
}

// кол-во друзей
func countFriends(m map[string][]string) (result map[string]int) {
	result = make(map[string]int)

	for key, value := range m {
		count := 0
		for i := 0; i < len(value); i++ {
			count++
		}
		result[key] = count
	}
	return result
}

// вывод информации
func outputCountFriends(m map[string]int) string {
	result := []string{}

	for key, value := range m {
		result = append(result, fmt.Sprintf("%s: %d", key, value))
	}
	slices.Sort(result)
	return fmt.Sprintf("Количество друзей: \n%s", strings.Join(result, "\n"))
}

// общие друзья
func commonFriends(m map[string][]string, a, b string) string {
	result := []string{}

	for _, v1 := range m[a] {
		for _, v2 := range m[b] {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}
	slices.Sort(result)
	return fmt.Sprintf("Общие друзья между пользователями %s и %s: %s.", a, b, strings.Join(result, ", "))
}

// самые популярные пользователи
func mostPopularUsers(m map[string][]string) string {
	result := []string{}
	maxCount := 0

	for _, friends := range m {
		if len(friends) > maxCount {
			maxCount = len(friends)
		}
	}
	for name, friends := range m {
		if maxCount == len(friends) {
			result = append(result, name)
		}
	}
	slices.Sort(result)

	return fmt.Sprintf("Наиболее популярные пользователи: %s (кол-во друзей: %d).", strings.Join(result, ", "), maxCount)
}
