package model

import (
	"errors"
	"fmt"
)

var (
	errID     = errors.New("wrong ID")
	errTitle  = errors.New("title can't be empty")
	errAuthor = errors.New("author can't be empty")
)

type Book struct {
	ID     int
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("ID: %d, Название: %s, Автор: %s", b.ID, b.Title, b.Author)
}

func NewBook(id int, title string, author string) (Book, error) {
	if id <= 0 {
		return Book{}, errID
	}
	if title == "" {
		return Book{}, errTitle
	}
	if author == "" {
		return Book{}, errAuthor
	}
	return Book{
		ID:     id,
		Title:  title,
		Author: author,
	}, nil
}
