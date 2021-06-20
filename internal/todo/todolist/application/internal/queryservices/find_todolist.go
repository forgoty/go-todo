package queryservices

import (
	"github.com/foroto/go-todo/internal/todo/todolist/domain/model/aggregates"
	"github.com/google/uuid"
)

type FindToDoListByIdQueryHandler struct {
	toDoListRepo aggregates.ToDoListRepository
}

func NewFindToDoListByIdQueryHandler(repo *aggregates.ToDoListRepository) FindToDoListByIdQueryHandler {
	return FindToDoListByIdQueryHandler{
		toDoListRepo: *repo,
	}
}

func (c *FindToDoListByIdQueryHandler) Handle(id string) (*aggregates.ToDoList, error) {
	uuidStr, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return c.toDoListRepo.FindById(uuidStr)
}
