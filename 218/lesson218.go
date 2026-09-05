package main

import (
	"fmt"
	"lesson/218/worker"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("worker error: %s", err)
	}
}

func run() error {
	w := worker.New()
	if err := w.DoWork("file.exe"); err != nil {
		return fmt.Errorf("do work: %w", err)
	}
	return nil
}
