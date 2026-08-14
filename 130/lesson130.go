package main

import "fmt"

// находим индекс нужного символа в строке
// и значение ok: true, если символ найден, иначе false
func Index(str string, a rune) (index int, ok bool) {
	for i, c := range str {
		if c == a {
			return i, true
		}
	}
	// если не нашли - возвращаем -1 и false, чтобы было понятнее
	return -1, false
}

func main() {
	tests := []struct {
		s    string
		r    rune
		want int
		ok   bool
	}{
		{"hello", 'e', 1, true},
		{"golang", 'g', 0, true},   // первая 'g'
		{"golang", 'z', -1, false}, // нет такого символа
		{"Привет", 'и', 3, true},   // кириллица — range корректно работает с UTF‑8
		{"", 'a', -1, false},       // пустая строка
		{"aaa", 'a', 0, true},      // повторяющиеся символы — берём первый
	}

	for _, t := range tests {
		idx, ok := Index(t.s, t.r)
		fmt.Printf("String: %q, rune: %q -> index: %d, ok: %v (expected: %d, %v)\n",
			t.s, t.r, idx, ok, t.want, t.ok)
	}
}
