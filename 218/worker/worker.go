package worker

type Worker struct{}

func New() *Worker {
	return &Worker{}
}

func (Worker) DoWork(file string) error {
	return NewFileNotFoundError(file)
}
