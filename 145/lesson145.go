package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	m1 := map[string]int{
		"Linus Torvalds":  5,
		"James Gosling":   3,
		"Tim Berners-Lee": 4,
	}
	m2 := map[string]int{
		"Mitchel Resnick":   0,
		"Linus Torvalds":    5,
		"Donald Knuth":      3,
		"Tim Berners-Lee":   5,
		"Bjarne Stroustrup": 5,
	}
	m3 := map[string]int{
		"Candidate1": 0,
		"Candidate2": 0,
		"Candidate3": 0,
	}
	m4 := map[string]int{}

	fmt.Print(countVotes(m1))
	fmt.Print(countVotes(m2))
	fmt.Print(countVotes(m3))
	fmt.Print(countVotes(m4))
}

func countVotes(votes map[string]int) string {
	if votes == nil {
		return fmt.Sprintln("Ошибка, пустой ввод.")
	}
	allNames := ""
	sumVotes := 0
	maxCount := 0
	for key, value := range votes {
		allNames += key
		sumVotes += value
		if value > maxCount {
			maxCount = value
		}
	}
	if allNames == "" {
		return fmt.Sprintln("Кандидаты потерялись.")
	}
	if sumVotes == 0 {
		return fmt.Sprintln("Все голоса похищены!")
	}
	names := []string{}
	for key, value := range votes {
		if value == maxCount {
			names = append(names, key)
		}
	}
	slices.Sort(names)
	return fmt.Sprintln(strings.Join(names, ", "))
}
