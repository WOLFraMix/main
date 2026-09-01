package main

import "fmt"

func main() {
	myDB := MyDB{}
	logger := logger{}

	run(myDB, logger)
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

type DB interface {
	Reader
	Writer
}

type MyDB struct {
}

func (m MyDB) Write(str string) {
	fmt.Printf("Записали %q\n", str)
}

func (m MyDB) Read() string {
	return "Вернули нужную строку"
}

func run(db DB, logger Writer) {
	db.Write("Hello")
	fmt.Println(db.Read())

	logger.Write("Лог")
}

type logger struct{}

func (l logger) Write(str string) {
	fmt.Printf("Залогировали %q\n", str)
}
