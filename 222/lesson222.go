package main

import (
	"errors"
	"lesson/222/networkerr"
	"lesson/222/worker"
	"log"
)

func main() {
	w := worker.New()
	if err := w.DoWork(); err != nil {
		var networkerr *networkerr.NetworkError
		switch {
		case errors.Is(err, worker.ErrNoInternet):
			log.Fatalf("check your internet connection")
		case errors.As(err, &networkerr):
			log.Fatalf("network worker error (code: %d): %s", networkerr.Code, networkerr)
		default:
			log.Fatalf("unknown worker error: %s", err)
		}
	}
}
