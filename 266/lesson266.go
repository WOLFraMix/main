package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	phoneBook, err := NewSimpleMap(1)
	if err != nil {
		panic(fmt.Sprintf("failed to create map: %v", err))
	}

	fmt.Printf("Начальное количество бакетов: %d\n", len(phoneBook.buckets))

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
	phoneBook.Put("Джеймс", "00700")

	fmt.Printf("Количество бакетов после вставки значений: %d\n", len(phoneBook.buckets))

	name := "Иван"
	if num, ok := phoneBook.Get(name); ok {
		fmt.Printf("%s: %s\n", name, num)
	} else {
		fmt.Printf("Пользователя %s не существует\n", name)
	}

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
	size    int
	count   int
}

// NewSimpleMap создает hash-таблицу
func NewSimpleMap(size int) (*SimpleMap, error) {
	if size <= 0 {
		return nil, errors.New("size must be positive")
	}
	return &SimpleMap{
		buckets: make([][]Entry, size),
		size:    size,
		count:   0,
	}, nil
}

// Метод Put кладет новое значение или меняет значение,
// если такой ключ уже существует.
func (m *SimpleMap) Put(key, value string) {
	h := hash(key)
	idx := m.bucketIndex(h)
	bucket := m.buckets[idx]

	// Ищем, есть ли уже такой ключ
	for i, entry := range bucket {
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
	if m.count > m.size*10 {
		m.Resize()
	}
}

// Метод Resize расширяет hash-таблицу
func (m *SimpleMap) Resize() {
	newSize := m.size * 2
	newBuckets := make([][]Entry, newSize)
	for _, bucket := range m.buckets {
		for _, entry := range bucket {
			idx := entry.Hash % uint(newSize)
			newBuckets[idx] = append(newBuckets[idx], entry)
		}
	}
	m.buckets = newBuckets
	m.size = newSize
}

// Метод Get возвращает данные, с подтверждением их существования.
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

// Len возвращает количество значений в SimpleMap.
func (m *SimpleMap) Len() int {
	return m.count
}

// Has проверяет, существует ли значение в SimpleMap.
func (m *SimpleMap) Has(key string) bool {
	_, ok := m.Get(key)
	return ok
}

// Delete удаляет значение по ключу
func (m *SimpleMap) Delete(key string) {
	h := hash(key)
	idx := m.bucketIndex(h)

	bucket := m.buckets[idx]
	for i := range bucket {
		if bucket[i].Hash == h && bucket[i].Key == key {
			// Удаляем элемент, сохраняя плотность слайса.
			lastIdx := len(bucket) - 1
			bucket[i] = bucket[lastIdx]
			// Присваиваем нулевое значение последнему элементу перед тем как обрежем,
			// чтобы GC мог собрать строки, если они больше нигде не используются.
			bucket[lastIdx] = Entry{} // Нулевое значение ставим для того чтобы GC собрал мусор
			m.buckets[idx] = bucket[:lastIdx]
			m.count--
			return
		}
	}
}

// Keys возвращает список ключей в SimpleMap
func (m *SimpleMap) Keys() []string {
	keys := make([]string, 0, m.count)
	for _, bucket := range m.buckets {
		for _, entry := range bucket {
			keys = append(keys, entry.Key)
		}
	}
	return keys
}

// Values возвращает список значений в SimpleMap
func (m *SimpleMap) Values() []string {
	values := make([]string, 0, m.count)
	for _, bucket := range m.buckets {
		for _, entry := range bucket {
			values = append(values, entry.Value)
		}
	}
	return values
}

// Clear очищает все значения в SimpleMap
func (m *SimpleMap) Clear() {
	// Очищаем каждый бакет, но сохраняем capacity массива бакетов,
	// это позволяет переиспользовать память при последующих вставках.
	for i := range m.buckets {
		m.buckets[i] = nil // позволяем GC собрать Entries
	}
	m.count = 0
}

// bucketIndex дает индекс бакета по хешу
func (m *SimpleMap) bucketIndex(h uint) uint {
	return h % uint(m.size)
}

// hash функция для получения хеша от строки
func hash(s string) uint {
	var h uint = 0
	for _, r := range s {
		h = h*31 + uint(r)
	}
	return h
}
