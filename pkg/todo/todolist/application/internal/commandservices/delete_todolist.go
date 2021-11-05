package commandservices

import (
	"github.com/forgoty/go-todo/pkg/todo/todolist/domain/model/aggregates"
	"github.com/google/uuid"
)

type DeleteToDoListCommandHandler struct {
	toDoListRepo aggregates.ToDoListRepository
}

func NewDeleteToDoListCommandHandler(repo *aggregates.ToDoListRepository) DeleteToDoListCommandHandler {
	return DeleteToDoListCommandHandler{
		toDoListRepo: *repo,
	}
}

func (c *DeleteToDoListCommandHandler) Handle(id string) error {
	uuidStr, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = c.toDoListRepo.DeleteById(uuidStr)
	if err != nil {
		return err
	}
	return nil
}
