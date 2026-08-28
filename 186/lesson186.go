package main

import (
	"errors"
	"fmt"
)

var (
	ErrTagExists    = errors.New("tag already exists")
	ErrTagNotExists = errors.New("tag not exists")
)

type TagManager struct {
	tags map[string]struct{}
}

func NewTagManager() *TagManager {
	return &TagManager{
		tags: make(map[string]struct{}),
	}
}

func (tm *TagManager) AddTag(tag string) error {
	if _, exists := tm.tags[tag]; exists {
		return ErrTagExists
	}
	tm.tags[tag] = struct{}{}
	return nil
}

func (tm *TagManager) RemoveTag(tag string) error {
	if _, exists := tm.tags[tag]; !exists {
		return ErrTagNotExists
	}
	delete(tm.tags, tag)
	return nil
}

func (tm *TagManager) TagExists(tag string) bool {
	if _, exists := tm.tags[tag]; exists {
		return true
	}
	return false
}

func (tm *TagManager) ListTags() []string {
	list := make([]string, 0, len(tm.tags))
	for tag := range tm.tags {
		list = append(list, tag)
	}
	return list
}

func main() {
	tm := NewTagManager()

	// Добавление тегов
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err)
	}
	if err := tm.AddTag("programming"); err != nil {
		fmt.Println(err)
	}
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег уже существует
	}

	// Проверка существования тегов
	fmt.Println("Тег 'golang' существует:", tm.TagExists("golang")) // true
	fmt.Println("Тег 'python' существует:", tm.TagExists("python")) // false

	// Список тегов
	fmt.Println("Current tags:", tm.ListTags()) // [golang programming]

	// Удаление тегов
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err)
	}
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег не существует
	}

	// Список тегов после удаления
	fmt.Println("Current tags after removal:", tm.ListTags()) // [programming]
}
