package main

import (
	"fmt"
)

func main() {
	phoneBook := NewSimpleMap(10)
	phoneBook.Put("Анна", "8937593272")
	phoneBook.Put("Павел", "9837593482")
	phoneBook.Put("Мария", "2857929872")
	phoneBook.Put("Елена", "9048539849")
	phoneBook.Put("Софья", "5798237599")
	phoneBook.Put("Дарья", "2981739812")
	phoneBook.Put("Кирилл", "3733733397")
	phoneBook.Put("Иван", "7895378888")
	phoneBook.Put("Виктор", "9923782474")
	phoneBook.Put("Алексей", "1192849002")

	name := "Иван"
	if num, ok := phoneBook.Get(name); ok {
		fmt.Printf("%s: %s\n", name, num)
	} else {
		fmt.Printf("Пользователя %s не существует\n", name)
	}
}

type Entry struct {
	Key   string
	Value string
}

type SimpleMap struct {
	buckets [][]Entry
	size    int
}

func NewSimpleMap(size int) *SimpleMap {
	return &SimpleMap{
		buckets: make([][]Entry, size),
		size:    size,
	}
}

// Функция кладет новое значение или меняет значение, если такой ключ уже существует
func (m *SimpleMap) Put(key, value string) {
	idx := hash(key) % uint(m.size)
	// Ищем, есть ли уже такой ключ
	for i, entry := range m.buckets[idx] {
		if entry.Key == key {
			m.buckets[idx][i].Value = value // обновляем
			return
		}
	}
	// Если нет - добавляем новую запись
	m.buckets[idx] = append(m.buckets[idx], Entry{Key: key, Value: value})
}

// Функция возвращает данные с подтверждением, были они или нет
func (m *SimpleMap) Get(key string) (string, bool) {
	idx := hash(key) % uint(m.size)
	for _, entry := range m.buckets[idx] {
		if entry.Key == key {
			return entry.Value, true
		}
	}
	return "", false
}

func hash(s string) uint {
	var h uint = 0
	for _, r := range s {
		h = h*31 + uint(r)
	}
	return h
}
