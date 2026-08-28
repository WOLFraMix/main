package project

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Status string

const (
	StatusActive Status = "active"
	StatusClosed Status = "closed"
)

type Project struct {
	ID    uuid.UUID
	Name  string
	Tasks []Task
}

func New(id uuid.UUID, name string) (*Project, error) {
	if len(name) == 0 {
		return nil, errors.New("project need name")
	}
	return &Project{
		ID:    id,
		Name:  name,
		Tasks: make([]Task, 0),
	}, nil
}

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      Status
}

func NewTask(id uuid.UUID, head string, description string) (*Task, error) {
	if len(head) == 0 || len(description) == 0 {
		return nil, errors.New("task needs a title and description")
	}
	return &Task{
		ID:          id,
		Title:       head,
		Description: description,
		Status:      StatusActive,
	}, nil
}

func (p *Project) AddTask(task Task) error {
	for _, existingTasks := range p.Tasks {
		if existingTasks.ID == task.ID {
			return errors.New("task ID already exists")
		}
	}
	p.Tasks = append(p.Tasks, task)
	return nil
}

func (p *Project) UpdateTask(task Task) error {
	for i, existingTasks := range p.Tasks {
		if existingTasks.ID == task.ID {
			p.Tasks[i] = task
			return nil
		}
	}
	return errors.New("task ID doesn't exist")
}

func (t *Task) Close() error {
	if t.Status == StatusActive {
		t.Status = StatusClosed
		return nil
	}
	return errors.New("task status already closed")
}

func (t *Task) UpdateDescription(description string) error {
	if t.Status == StatusActive {
		t.Description = description
		return nil
	}
	return errors.New("update task failure")
}

func (p *Project) FilterTasksByStatus(status Status) []Task {
	var list []Task
	for _, task := range p.Tasks {
		if task.Status == status {
			list = append(list, task)
		}
	}
	return list
}

func (p *Project) PrintInfo() {
	fmt.Printf("%+v", p)
}
