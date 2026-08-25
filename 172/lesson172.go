package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name   string
	Grades []int
}

func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range s.Grades {
		sum += float64(v)
	}
	result := sum / float64(len(s.Grades))
	return roundTo(result, 1)
}

// roundTo округление х до num знаков после запятой
func roundTo(x float64, num int) float64 {
	pow := math.Pow(10, float64(num))
	return math.Round(x*pow) / pow
}

func (s Student) Info() string {
	return fmt.Sprintf("Студент %s, средняя оценка: %.1f.", s.Name, s.AverageGrade())
}

func main() {
	s1 := Student{
		Name:   "Katya",
		Grades: []int{4, 5, 4, 5, 5, 3, 4, 5, 3, 5, 5},
	}

	fmt.Println(s1.AverageGrade())
	fmt.Println(s1.Info())
}
