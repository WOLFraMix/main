package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	s := make([]int, n)

	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		s[i] = v
	}

	s[a] = 1000
	fmt.Println(s)
	s = append(s, a, b)
	fmt.Println(s)
	var sum int
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println(s[len(s)-a:])
	fmt.Println(s[a:b])
}
