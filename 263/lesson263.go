package main

import (
	"fmt"
)

func main() {
	arr := [10]string{}
	arr[hashIndex("Анна", len(arr))] = "9827382738"
	arr[hashIndex("Павел", len(arr))] = "8728472821"
	arr[hashIndex("Мария", len(arr))] = "7273723232"
	arr[hashIndex("Елена", len(arr))] = "2348539849"

	fmt.Printf("Анна, hash: %d, индекс: %d, телефон: %s\n", hash("Анна"), hashIndex("Анна", len(arr)), arr[hashIndex("Анна", len(arr))])
	fmt.Printf("Павел, hash: %d, индекс: %d, телефон: %s\n", hash("Павел"), hashIndex("Павел", len(arr)), arr[hashIndex("Павел", len(arr))])
	fmt.Printf("Мария, hash: %d, индекс: %d, телефон: %s\n", hash("Мария"), hashIndex("Мария", len(arr)), arr[hashIndex("Мария", len(arr))])
	fmt.Printf("Елена, hash: %d, индекс: %d, телефон: %s\n", hash("Елена"), hashIndex("Елена", len(arr)), arr[hashIndex("Елена", len(arr))])
}

func hashIndex(s string, size int) int {
	return int(hash(s) % uint(size))
}

func hash(s string) uint {
	var h uint = 0
	for _, r := range s {
		h = h*31 + uint(r)
	}
	return h
}
