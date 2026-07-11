package main

import (
	"math"
	"testing"
)

// --- Тесты для calculate ---

func TestCalculate_Add(t *testing.T) {
	tests := []struct {
		name  string
		a     float64
		b     float64
		want  float64
	}{
		{"простое сложение", 1, 2, 3},
		{"отрицательные числа", -1, -2, -3},
		{"смешанные", -1, 2, 1},
		{"ноль", 0, 0, 0},
		{"большие числа", 1e10, 2e10, 3e10},
		{"дробные", 1.5, 2.5, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.a, tt.b, "add")
			if err != nil {
				t.Errorf("calculate(%v, %v, \"add\") вернул ошибку: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Errorf("calculate(%v, %v, \"add\") = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculate_Subtract(t *testing.T) {
	tests := []struct {
		name  string
		a     float64
		b     float64
		want  float64
	}{
		{"простое вычитание", 1, 2, -1},
		{"отрицательные", -5, -3, -2},
		{"ноль", 0, 0, 0},
		{"дробные", 5.5, 2.5, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.a, tt.b, "subtract")
			if err != nil {
				t.Errorf("calculate(%v, %v, \"subtract\") вернул ошибку: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Errorf("calculate(%v, %v, \"subtract\") = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculate_Multiply(t *testing.T) {
	tests := []struct {
		name  string
		a     float64
		b     float64
		want  float64
	}{
		{"простое умножение", 3, 4, 12},
		{"с нулём", 5, 0, 0},
		{"отрицательные", -3, -4, 12},
		{"смешанные", -3, 4, -12},
		{"дробные", 2.5, 4.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.a, tt.b, "multiply")
			if err != nil {
				t.Errorf("calculate(%v, %v, \"multiply\") вернул ошибку: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Errorf("calculate(%v, %v, \"multiply\") = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculate_Divide(t *testing.T) {
	tests := []struct {
		name  string
		a     float64
		b     float64
		want  float64
	}{
		{"простое деление", 10, 2, 5},
		{"дробное", 7, 2, 3.5},
		{"отрицательное", -10, 2, -5},
		{"оба отрицательные", -10, -2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.a, tt.b, "divide")
			if err != nil {
				t.Errorf("calculate(%v, %v, \"divide\") вернул ошибку: %v", tt.a, tt.b, err)
			}
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("calculate(%v, %v, \"divide\") = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculate_DivideByZero(t *testing.T) {
	_, err := calculate(10, 0, "divide")
	if err == nil {
		t.Error("calculate(10, 0, \"divide\") должна возвращать ошибку, но получила nil")
	}
}

func TestCalculate_UnknownOperation(t *testing.T) {
	_, err := calculate(1, 2, "modulo")
	if err == nil {
		t.Error("calculate(1, 2, \"modulo\") должна возвращать ошибку, но получила nil")
	}
}

func TestCalculate_AllOperations(t *testing.T) {
	a, b := 10.0, 3.0

	// Проверка всех операций
	got, err := calculate(a, b, "add")
	if err != nil || math.Abs(got-13.0) > 1e-9 {
		t.Errorf("add: got %v, want 13.0, err: %v", got, err)
	}

	got, err = calculate(a, b, "subtract")
	if err != nil || math.Abs(got-7.0) > 1e-9 {
		t.Errorf("subtract: got %v, want 7.0, err: %v", got, err)
	}

	got, err = calculate(a, b, "multiply")
	if err != nil || math.Abs(got-30.0) > 1e-9 {
		t.Errorf("multiply: got %v, want 30.0, err: %v", got, err)
	}

	got, err = calculate(a, b, "divide")
	if err != nil || math.Abs(got-10.0/3.0) > 1e-9 {
		t.Errorf("divide: got %v, want %.9f, err: %v", got, 10.0/3.0, err)
	}
}

func TestCalculate_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		op    string
		a     float64
		b     float64
	}{
		{"отрицательное число +", "add", -100, -200},
		{"отрицательное число -", "subtract", -100, -200},
		{"отрицательное число *", "multiply", -100, -200},
		{"отрицательное число /", "divide", -100, -200},
		{"очень маленькое число", "divide", 1e-10, 1e-10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := calculate(tt.a, tt.b, tt.op)
			if err != nil {
				t.Errorf("calculate(%v, %v, %q) unexpected error: %v", tt.a, tt.b, tt.op, err)
			}
		})
	}
}
