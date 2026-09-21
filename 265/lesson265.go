package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	phoneBook, err := NewSimpleMap(10)
	if err != nil {
		log.Fatalf("NewSimpleMap: %v", err)
	}
	phoneBook.Put("Анна", "11111")
	phoneBook.Put("Павел", "22222")
	phoneBook.Put("Мария", "33333")
	phoneBook.Put("Елена", "44444")
	phoneBook.Put("Софья", "55555")
	phoneBook.Put("Дарья", "66666")
	phoneBook.Put("Кирилл", "77777")
	phoneBook.Put("Иван", "88888")
	phoneBook.Put("Виктор", "99999")
	phoneBook.Put("Алексей", "00000")

	name := "Иван"
	if num, ok := phoneBook.Get(name); ok {
		fmt.Printf("%s: %s\n", name, num)
	} else {
		fmt.Printf("Пользователя %s не существует\n", name)
	}

	fmt.Println(phoneBook.Len())

	fmt.Println()

	if phoneBook.Has(name) {
		fmt.Printf("Пользователь %s существует\n", name)
	} else {
		fmt.Printf("Пользователь %s не существует\n", name)
	}

	fmt.Printf("Удаляем пользователя %s\n", name)
	phoneBook.Delete(name)

	if phoneBook.Has(name) {
		fmt.Printf("Пользователь %s существует\n", name)
	} else {
		fmt.Printf("Пользователь %s не существует\n", name)
	}

	fmt.Println()

	fmt.Printf(
		"Все ключи: [%s]\nВсе значения: [%s]\n",
		strings.Join(phoneBook.Keys(), ", "),
		strings.Join(phoneBook.Values(), ", "),
	)

	fmt.Println()

	fmt.Printf("Очищаем SimpleMap\n")
	phoneBook.Clear()

	fmt.Printf(
		"Все ключи: [%s]\nВсе значения: [%s]\n",
		strings.Join(phoneBook.Keys(), ", "),
		strings.Join(phoneBook.Values(), ", "),
	)

	fmt.Println()

	fmt.Println(phoneBook.Len())
}

// Entry хранит данные о значении в SimpleMap
type Entry struct {
	Key   string
	Value string
	Hash  uint
}

// SimpleMap простейшая hash-таблица
type SimpleMap struct {
	buckets [][]Entry
	size    uint
	count   uint
}

// NewSimpleMap создает hash-таблицу
func NewSimpleMap(size uint) (*SimpleMap, error) {
	if size <= 0 {
		return nil, errors.New("size must be positive")
	}
	return &SimpleMap{
		buckets: make([][]Entry, size),
		size:    size,
	}, nil
}

// Метод Put кладет новое значение или меняет значение,
// если такой ключ уже существует.
func (m *SimpleMap) Put(key, value string) {
	h := hash(key)
	idx := m.bucketIndex(h)
	// Ищем, есть ли уже такой ключ
	for i, entry := range m.buckets[idx] {
		if entry.Hash == h && entry.Key == key {
			m.buckets[idx][i].Value = value // обновляем
			return
		}
	}
	// Если нет - добавляем новую запись
	m.buckets[idx] = append(m.buckets[idx], Entry{
		Key:   key,
		Value: value,
		Hash:  h,
	})
	m.count++
}

// Метод Get возвращает данные,
// с подтверждением их существования.
func (m *SimpleMap) Get(key string) (string, bool) {
	h := hash(key)
	idx := m.bucketIndex(h)
	for _, entry := range m.buckets[idx] {
		if entry.Hash == h && entry.Key == key {
			return entry.Value, true
		}
	}
	return "", false
}

// hash функция для получения хэша от строки
func hash(s string) uint {
	var h uint = 0
	for _, r := range s {
		h = h*31 + uint(r)
	}
	return h
}

func (m *SimpleMap) Len() uint {
	return m.count
}

// Метод bucketIndex возвращает индекс
func (m *SimpleMap) bucketIndex(h uint) uint {
	return h % uint(m.size)
}

// Метод Has проверяет, есть ли значение в hash-таблице
func (m *SimpleMap) Has(key string) bool {
	h := hash(key)
	idx := m.bucketIndex(h)
	for _, entry := range m.buckets[idx] {
		if entry.Hash == h && entry.Key == key {
			return true
		}
	}

	return false
}

// Метод Delete удаляет значение по ключу
func (m *SimpleMap) Delete(key string) {
	h := hash(key)
	idx := m.bucketIndex(h)
	if m.Has(key) {
		for i, entry := range m.buckets[idx] {
			if entry.Hash == h && entry.Key == key {
				m.buckets[idx] = append(m.buckets[idx][:i], m.buckets[idx][i+1:]...)
				m.count--
			}
		}
	}
}

// Метод Keys возвращает все ключи
func (m *SimpleMap) Keys() []string {
	var result []string
	for _, v := range m.buckets {
		for _, k := range v {
			result = append(result, k.Key)
		}
	}
	return result
}

// Метод Values возвращает все значения
func (m *SimpleMap) Values() []string {
	var result []string
	for _, v := range m.buckets {
		for _, val := range v {
			result = append(result, val.Value)
		}
	}
	return result
}

// Метод Clear очищает hash-таблицу
func (m *SimpleMap) Clear() {
	m.buckets = nil
	m.count = 0
}
