package main

import "lesson/201/worker"

func main() {
	worker := worker.New()
	worker.CreateBooks()
	worker.PrintBook(5)
}
