package main

import "fmt"

// входит ли одна строка в другую
func isSubsequence(a, b string) bool {
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	// если в конце цикла
	// индекс i равен длине строки a,
	// значит, все символы из a
	// были найдены в b в правильном порядке
	return i == len(a)
}

func main() {
	tests := []struct {
		a      string
		b      string
		expect bool
	}{
		{"ace", "abcde", true},   // Все символы 'a', 'c', 'e' найдены в порядке следования
		{"aec", "abcde", false},  // Символы есть, но порядок нарушен ('e' встречается раньше 'c')
		{"abc", "ahbgdc", true},  // Пропущены лишние символы 'h', 'g', 'd', порядок сохранён
		{"axc", "ahbgdc", false}, // Символ 'x' отсутствует в строке b
		{"", "abc", true},        // Пустая строка всегда является подпоследовательностью
		{"abc", "", false},       // Непустая строка не может быть подпоследовательностью пустой
	}

	for _, t := range tests {
		result := isSubsequence(t.a, t.b)
		status := "✅"
		if result != t.expect {
			status = "❌"
		}
		fmt.Printf("%s isSubsequence(%q, %q) = %v (ожидалось %v)\n", status, t.a, t.b, result, t.expect)
	}
}
