package worker

import (
	"errors"
	"lesson/222/networkerr"
	"math/rand/v2"
)

var (
	ErrServiceUnavailable       = networkerr.New("service unavailable", 233)
	ErrNoInternet         error = errors.New("no internet")
)

type Worker struct {
}

func New() *Worker {
	return &Worker{}
}

func (Worker) DoWork() error {
	if rand.IntN(100) < 40 {
		return ErrServiceUnavailable
	}
	if rand.IntN(100) < 40 {
		return networkerr.NewWithErr(ErrNoInternet, 235)
	}
	return errors.New("file not found")
}
