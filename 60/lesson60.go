package main

type Day int

const (
	Monday Day = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func isWeekend(d Day) bool {
	return d == Saturday || d == Sunday
}

func main() {
	println(isWeekend(Monday))
}
