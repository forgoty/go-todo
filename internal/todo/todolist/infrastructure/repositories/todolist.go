package todolist

import (
	"fmt"
	"github.com/foroto/go-todo/internal/todo/todolist/domain/model/aggregates"
	"github.com/google/uuid"
)

type NotFoundError struct {
	ToDoListUUID uuid.UUID
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("To-Do List with ID '%s' not found", e.ToDoListUUID)
}

type InMemoryToDoListRepository struct {
	ToDoLists []aggregates.ToDoList
}

func NewInMemoryToDoListRepository() *InMemoryToDoListRepository {
	return &InMemoryToDoListRepository{}
}

func (repo *InMemoryToDoListRepository) FindAll() ([]aggregates.ToDoList, error) {
	return repo.ToDoLists, nil
}

func (repo *InMemoryToDoListRepository) FindById(id uuid.UUID) (*aggregates.ToDoList, error) {
	for _, tdl := range repo.ToDoLists {
		if tdl.Id == id {
			return &tdl, nil
		}
	}
	return nil, NotFoundError{ToDoListUUID: id}
}

func (repo *InMemoryToDoListRepository) Store(tdl aggregates.ToDoList) error {
	repo.ToDoLists = append(repo.ToDoLists, tdl)
	return nil
}

func (repo *InMemoryToDoListRepository) DeleteById(id uuid.UUID) error {
	tmpIndex := -1
	for index, tdl := range repo.ToDoLists {
		if tdl.Id == id {
			tmpIndex = index
			break
		}
	}
	if tmpIndex < 0 {
		return NotFoundError{ToDoListUUID: id}
	}
	repo.ToDoLists = append(repo.ToDoLists[:tmpIndex], repo.ToDoLists[tmpIndex+1:]...)
	return nil
}
