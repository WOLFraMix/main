package main

import (
	"math"
	"testing"
)

// --- Тесты для логики PetBattle ---

type petResult struct {
	winner string // "cats", "dogs", "draw"
	score  string // формат "X:Y" или ""
}

func petBattleLogic(cats int, dogs int) petResult {
	if cats > dogs {
		return petResult{winner: "cats", score: "cats"}
	} else if cats < dogs {
		return petResult{winner: "dogs", score: "dogs"}
	}
	return petResult{winner: "draw", score: "draw"}
}

func TestPetBattleLogic(t *testing.T) {
	tests := []struct {
		name   string
		cats   int
		dogs   int
		expect petResult
	}{
		{"коты побеждают", 3, 2, petResult{winner: "cats"}},
		{"собаки побеждают", 1, 4, petResult{winner: "dogs"}},
		{"ничья", 5, 5, petResult{winner: "draw"}},
		{"коты побеждают 0:1", 0, 1, petResult{winner: "dogs"}},
		{"коты побеждают 10:0", 10, 0, petResult{winner: "cats"}},
		{"ничья 0:0", 0, 0, petResult{winner: "draw"}},
		{"большие числа", 100, 99, petResult{winner: "cats"}},
		{"отрицательные коты", -1, 0, petResult{winner: "dogs"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := petBattleLogic(tt.cats, tt.dogs)
			if result.winner != tt.expect.winner {
				t.Errorf("petBattleLogic(%d, %d).winner = %q, want %q",
					tt.cats, tt.dogs, result.winner, tt.expect.winner)
			}
		})
	}
}

// --- Тесты для логики printNumberInfo ---

type numberInfo struct {
	sign     string // "zero", "negative", "positive"
	parity   string // "even", "odd"
	sqrtType string // "perfect", "irrational", "n/a" (для 0 и отрицательных)
	sqrtVal  float64
}

func numberInfoLogic(num int) numberInfo {
	info := numberInfo{}

	switch {
	case num == 0:
		info.sign = "zero"
	case num < 0:
		info.sign = "negative"
	case num > 0:
		info.sign = "positive"
	}

	switch {
	case num%2 == 0:
		info.parity = "even"
	case num%2 != 0:
		info.parity = "odd"
	}

	if num > 0 {
		sqrtNum := math.Sqrt(float64(num))
		if sqrtNum == math.Floor(sqrtNum) {
			info.sqrtType = "perfect"
			info.sqrtVal = sqrtNum
		} else {
			info.sqrtType = "irrational"
			info.sqrtVal = sqrtNum
		}
	} else {
		info.sqrtType = "n/a"
	}

	return info
}

func TestNumberInfoLogic(t *testing.T) {
	tests := []struct {
		name   string
		num    int
		expect numberInfo
	}{
		{"отрицательное нечётное", -4, numberInfo{sign: "negative", parity: "even", sqrtType: "n/a"}},
		{"положительное нечётное", 5, numberInfo{sign: "positive", parity: "odd", sqrtType: "irrational", sqrtVal: math.Sqrt(5)}},
		{"положительное чётное, полный корень", 16, numberInfo{sign: "positive", parity: "even", sqrtType: "perfect", sqrtVal: 4}},
		{"ноль", 0, numberInfo{sign: "zero", parity: "even", sqrtType: "n/a"}},
		{"отрицательное нечётное", -3, numberInfo{sign: "negative", parity: "odd", sqrtType: "n/a"}},
		{"полный корень 25", 25, numberInfo{sign: "positive", parity: "odd", sqrtType: "perfect", sqrtVal: 5}},
		{"полный корень 1", 1, numberInfo{sign: "positive", parity: "odd", sqrtType: "perfect", sqrtVal: 1}},
		{"большое число", 100, numberInfo{sign: "positive", parity: "even", sqrtType: "perfect", sqrtVal: 10}},
		{"не полный корень 2", 2, numberInfo{sign: "positive", parity: "even", sqrtType: "irrational", sqrtVal: math.Sqrt(2)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numberInfoLogic(tt.num)
			if result.sign != tt.expect.sign {
				t.Errorf("numberInfoLogic(%d).sign = %q, want %q", tt.num, result.sign, tt.expect.sign)
			}
			if result.parity != tt.expect.parity {
				t.Errorf("numberInfoLogic(%d).parity = %q, want %q", tt.num, result.parity, tt.expect.parity)
			}
			if result.sqrtType != tt.expect.sqrtType {
				t.Errorf("numberInfoLogic(%d).sqrtType = %q, want %q", tt.num, result.sqrtType, tt.expect.sqrtType)
			}
			if tt.expect.sqrtType != "n/a" && result.sqrtVal != tt.expect.sqrtVal {
				t.Errorf("numberInfoLogic(%d).sqrtVal = %v, want %v", tt.num, result.sqrtVal, tt.expect.sqrtVal)
			}
		})
	}
}

func TestNumberInfoLogic_EdgeCases(t *testing.T) {
	tests := []struct {
		num    int
		sign   string
		parity string
	}{
		{-100, "negative", "even"},
		{-99, "negative", "odd"},
		{1, "positive", "odd"},
		{2, "positive", "even"},
		{2147483647, "positive", "odd"},
		{-2147483648, "negative", "even"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := numberInfoLogic(tt.num)
			if result.sign != tt.sign {
				t.Errorf("numberInfoLogic(%d).sign = %q, want %q", tt.num, result.sign, tt.sign)
			}
			if result.parity != tt.parity {
				t.Errorf("numberInfoLogic(%d).parity = %q, want %q", tt.num, result.parity, tt.parity)
			}
		})
	}
}
