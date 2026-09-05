package worker

import "fmt"

type FileNotFoundError struct {
	File string
}

func NewFileNotFoundError(file string) *FileNotFoundError {
	return &FileNotFoundError{
		File: file,
	}
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("file %s is not found", e.File)
}
