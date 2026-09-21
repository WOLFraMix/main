package main

import "fmt"

func main() {
	anna := hash("anna")
	fmt.Println(anna)

	arr := [10]string{}
	arr[hashIndex("anna", len(arr))] = "здесь записи об Анне"

	name := "anna"
	idx := hashIndex(name, len(arr))
	fmt.Println(arr[idx])
}

func hash(s string) uint {
	var hash uint
	for _, r := range s {
		hash += uint(r)
	}
	return hash
}

func hashIndex(s string, size int) int {
	return int(hash(s)) % size
}
