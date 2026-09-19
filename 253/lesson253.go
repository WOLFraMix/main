package main

import "math"

type Range struct {
	From int `json:"from"` // начальный индекс
	To   int `json:"to"`   // конечный индекс (включительно)
}

// FindMaxAvgWindows находит окна с максимальной средней температурой.
func FindMaxAvgWindows(temps []float64, k int) []Range {
	if k <= 0 || len(temps) < k {
		return nil
	}

	var sum float64
	var avg float64
	var maxAvg float64

	for i := 0; i < k; i++ {
		sum += temps[i]
	}
	avg = sum / float64(k)
	avg = math.Round(avg*10) / 10
	maxAvg = avg

	j := 0
	for i := k; i < len(temps); i++ {
		sum = sum + temps[i] - temps[j]
		j++
		avg = sum / float64(k)
		avg = math.Round(avg*10) / 10
		if avg > maxAvg {
			maxAvg = avg
		}
	}

	var result []Range
	sum = 0
	j = 0

	for i := 0; i < k; i++ {
		sum += temps[i]
	}
	avg = sum / float64(k)
	avg = math.Round(avg*10) / 10
	if avg == maxAvg {
		result = append(result, Range{From: 0, To: k - 1})
	}

	for i := k; i < len(temps); i++ {
		sum = sum + temps[i] - temps[j]
		j++
		avg = sum / float64(k)
		avg = math.Round(avg*10) / 10
		if avg == maxAvg {
			result = append(result, Range{From: j, To: i})
		}
	}

	return result
}
