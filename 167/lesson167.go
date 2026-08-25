package main

import "fmt"

type Event struct {
	Title    string // название события
	Date     string // дата события
	Location string // место проведения события
}

func main() {
	fmt.Printf("%+v\n", createGoEvent())
}

func createGoEvent() (result Event) {
	result = Event{
		Title:    "День рождения Golang",
		Date:     "10 ноября 2009",
		Location: "GoogleLand",
	}
	return
}
