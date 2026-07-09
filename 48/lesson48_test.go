package main

import (
	"fmt"
	"strings"
	"testing"
)

// --- Тесты для sum ---

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		n1       int
		n2       int
		expected int
	}{
		{"два положительных", 2, 5, 7},
		{"два отрицательных", -3, -4, -7},
		{"положительное и отрицательное", 10, -5, 5},
		{"ноль и число", 0, 100, 100},
		{"два нуля", 0, 0, 0},
		{"большие числа", 1000000, 2000000, 3000000},
		{"отрицательное и положительное", -10, 3, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sum(tt.n1, tt.n2)
			if result != tt.expected {
				t.Errorf("sum(%d, %d) = %d, want %d", tt.n1, tt.n2, result, tt.expected)
			}
		})
	}
}

func TestSum_Associativity(t *testing.T) {
	a, b, c := 10, 20, 30
	// (a + b) + c == a + (b + c)
	left := sum(sum(a, b), c)
	right := sum(a, sum(b, c))
	if left != right {
		t.Errorf("Ассоциативность нарушена: (%d+%d)+%d = %d, %d+(%d+%d) = %d",
			a, b, c, left, a, b, c, right)
	}
}

func TestSum_Commutativity(t *testing.T) {
	a, b := 42, 13
	if sum(a, b) != sum(b, a) {
		t.Errorf("Коммутативность нарушена: sum(%d, %d) = %d, sum(%d, %d) = %d",
			a, b, sum(a, b), b, a, sum(b, a))
	}
}

// --- Тесты для getFullLength ---

func TestGetFullLength(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		expectedLen int
		expectedRun int
	}{
		{"пустая строка", "", 0, 0},
		{"ASCII строка", "Hello", 5, 5},
		{"русский текст", "Привет, друг!", 23, 13},
		{"один символ", "A", 1, 1},
		{"один русский символ", "Я", 2, 1},
		{"эмодзи", "😀", 4, 1},
		{"смешанный текст", "Hi! Привет", 16, 10},
		{"пробелы", "   ", 3, 3},
		{"цифры", "12345", 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, runes := getFullLength(tt.str)
			if bytes != tt.expectedLen {
				t.Errorf("getFullLength(%q) bytes = %d, want %d", tt.str, bytes, tt.expectedLen)
			}
			if runes != tt.expectedRun {
				t.Errorf("getFullLength(%q) runes = %d, want %d", tt.str, runes, tt.expectedRun)
			}
		})
	}
}

func TestGetFullLength_Properties(t *testing.T) {
	tests := []string{"", "a", "привет", "Hello World", "日本語", "🎉🎊"}

	for _, str := range tests {
		t.Run(str, func(t *testing.T) {
			bytes, runes := getFullLength(str)
			// Количество байт всегда >= количества рун
			if bytes < runes {
				t.Errorf("bytes (%d) < runes (%d) для %q", bytes, runes, str)
			}
			// Для ASCII байты == руны
			allASCII := true
			for _, r := range str {
				if r > 127 {
					allASCII = false
					break
				}
			}
			if allASCII && bytes != runes {
				t.Errorf("Для ASCII строки bytes (%d) != runes (%d)", bytes, runes)
			}
		})
	}
}

// --- Тесты для generateCompliment ---

var validCompliments = []string{
	"Ты вызываешь восторг, %s!",
	"У тебя потрясающая улыбка, %s!",
	"Ты вдохновляешь, %s!",
}

func TestGenerateCompliment(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"имя на русском", "Катя"},
		{"имя на английском", "John"},
		{"однобуквенное имя", "A"},
		{"пустое имя", ""},
		{"имя с пробелами", "Mary Jane"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateCompliment(tt.name)

			// Результат не пустой
			if result == "" {
				t.Error("generateCompliment() вернул пустую строку")
			}

			// Результат содержит имя
			if tt.name != "" && !strings.Contains(result, tt.name) {
				t.Errorf("generateCompliment(%q) = %q не содержит имя", tt.name, result)
			}

			// Результат является одним из допустимых комплиментов
			isValid := false
			for _, tmpl := range validCompliments {
				expected := fmt.Sprintf(tmpl, tt.name)
				if result == expected {
					isValid = true
					break
				}
			}
			if !isValid {
				t.Errorf("generateCompliment(%q) = %q, не является одним из допустимых комплиментов", tt.name, result)
			}
		})
	}
}

func TestGenerateCompliment_Randomness(t *testing.T) {
	results := make(map[string]bool)

	// Вызываем много раз, чтобы собрать разные результаты
	for i := 0; i < 100; i++ {
		result := generateCompliment("Test")
		results[result] = true
	}

	// Поскольку rand.Seed не вызывается в тестах, результат может быть одинаковым.
	// Но проверяем, что результат всегда один из трёх вариантов.
	if len(results) > 3 {
		t.Errorf("Получено %d уникальных результатов, максимум 3", len(results))
	}

	// Все результаты должны быть валидными
	for result := range results {
		isValid := false
		for _, tmpl := range validCompliments {
			expected := fmt.Sprintf(tmpl, "Test")
			if result == expected {
				isValid = true
				break
			}
		}
		if !isValid {
			t.Errorf("Невалидный результат: %q", result)
		}
	}
}
