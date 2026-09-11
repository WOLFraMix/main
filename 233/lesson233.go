package main

import (
	"fmt"
	"os"
	"slices"
)

type Post struct {
	ID int `json:"id"`
}

func MergeFeeds(a, b []Post) []Post {
	result := slices.Clone(a)
	result = append(result, b...)
	return result
}

func Test1() error {
	a := []Post{{ID: 1}, {ID: 2}, {ID: 3}}
	aClone := []Post{{ID: 1}, {ID: 2}, {ID: 3}}

	b := []Post{{ID: 4}, {ID: 5}}
	bClone := []Post{{ID: 4}, {ID: 5}}

	result := MergeFeeds(a, b)
	result[1].ID = 100
	expected := []Post{{ID: 1}, {ID: 100}, {ID: 3}, {ID: 4}, {ID: 5}}

	// Проверяем, что a не изменился
	if !EqualSlices(a, aClone) {
		return fmt.Errorf("[Test1]: Исходный слайс 'a' был изменен.\nБыло: %v\nСтало: %v", aClone, a)
	}
	// Проверяем, что b не изменился
	if !EqualSlices(b, bClone) {
		return fmt.Errorf("[Test1]: Исходный слайс 'b' был изменен.\nБыло: %v\nСтало: %v", bClone, b)
	}
	// Проверяем результат
	if !EqualSlices(result, expected) {
		return fmt.Errorf("[Test1]: Результат неверный.\nОжидалось: %v\nПолучили: %v", expected, result)
	}
	return nil
}

func Test2() error {
	a := make([]Post, 2, 10)
	a[0] = Post{ID: 1}
	a[1] = Post{ID: 2}
	aClone := make([]Post, 2, 10)
	aClone[0] = Post{ID: 1}
	aClone[1] = Post{ID: 2}

	b := []Post{{ID: 3}, {ID: 4}, {ID: 5}}
	bClone := []Post{{ID: 3}, {ID: 4}, {ID: 5}}

	result := MergeFeeds(a, b)
	result[1].ID = 100
	expected := []Post{{ID: 1}, {ID: 100}, {ID: 3}, {ID: 4}, {ID: 5}}

	// Проверяем, что a не изменился
	if !EqualSlices(a, aClone) {
		return fmt.Errorf("[Test2]: Исходный слайс 'a' был изменен.\nБыло: %v\nСтало: %v", aClone, a)
	}
	// Проверяем, что b не изменился
	if !EqualSlices(b, bClone) {
		return fmt.Errorf("[Test2]: Исходный слайс 'b' был изменен.\nБыло: %v\nСтало: %v", bClone, b)
	}
	// Проверяем результат
	if !EqualSlices(result, expected) {
		return fmt.Errorf("[Test2]: Результат неверный.\nОжидалось: %v\nПолучили: %v", expected, result)
	}
	return nil
}

func EqualSlices(a, b []Post) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].ID != b[i].ID {
			return false
		}
	}
	return true
}

func main() {
	if err := Test1(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	if err := Test2(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Все тесты пройдены.")
}
