package main

import "fmt"

const (
	_  = iota             // игнорируем начальное значение
	KB = 1 << (10 * iota) // 1 << (10*1) = 1024
	MB = 1 << (10 * iota) // 1 << (10*2) = 1048576
	GB = 1 << (10 * iota) // 1 << (10*3) = 1073741824
	TB = 1 << (10 * iota) // 1 << (10*4) = 1099511627776
)

// «сдвинуть 1 на 10 бит влево», и это равно 2 в степени 10 и т.д.
func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}
