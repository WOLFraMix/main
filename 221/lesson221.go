package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrSimple = errors.New("simple error")
)

func main() {
	if err := fn1(); err != nil {
		if errors.Is(err, ErrSimple) {
			log.Fatalf("simple error occurred")
		} else {
			log.Fatalf("unknown error")
		}
	}
}

func fn1() error {
	return fmt.Errorf("fn2: %w", fn2())
}

func fn2() error {
	return ErrSimple
}
