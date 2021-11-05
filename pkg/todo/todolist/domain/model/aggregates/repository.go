package aggregates

import (
	"github.com/google/uuid"
)

type ToDoListRepository interface {
	FindAll() ([]ToDoList, error)
	FindById(id uuid.UUID) (*ToDoList, error)
	Store(tdl ToDoList) error
	DeleteById(id uuid.UUID) error
}
