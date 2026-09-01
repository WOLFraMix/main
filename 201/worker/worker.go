package worker

import (
	"fmt"
	"lesson/201/model"
	"lesson/201/repository"
)

type Worker struct {
	repo *repository.BookInMemoryRepository
}

func New() *Worker {
	return &Worker{
		repo: repository.NewBookInMemoryRepository(),
	}
}

func (w Worker) CreateBooks() error {
	b1, err := model.NewBook(5, "Война и Мир", "Лев Толстой")
	if err != nil {
		return fmt.Errorf("create new book %w", err)
	}
	if err := w.repo.Add(b1); err != nil {
		return fmt.Errorf("add book with ID %d in repository: %w", b1.ID, err)
	}

	b2, err := model.NewBook(10, "Преступление и наказание", "Фёдор Достоевский")
	if err != nil {
		return fmt.Errorf("create new book %w", err)
	}
	if err := w.repo.Add(b2); err != nil {
		return fmt.Errorf("add book with ID %d in repository: %w", b2.ID, err)
	}
	return nil
}

func (w Worker) PrintBook(ID int) error {
	b, err := w.repo.ByID(ID)
	if err != nil {
		return fmt.Errorf("get book by ID %d from repository: %w", ID, err)
	}
	fmt.Printf("%+v\n", b)
	return nil
}
