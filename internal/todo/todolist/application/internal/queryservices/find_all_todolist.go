package queryservices

import (
	"github.com/foroto/go-todo/internal/todo/todolist/domain/model/aggregates"
)

type FindAllToDoListByIdQueryHandler struct {
	toDoListRepo aggregates.ToDoListRepository
}

func NewFindAllToDoListByIdQueryHandler(repo *aggregates.ToDoListRepository) FindAllToDoListByIdQueryHandler {
	return FindAllToDoListByIdQueryHandler{
		toDoListRepo: *repo,
	}
}

func (c *FindAllToDoListByIdQueryHandler) Handle() ([]aggregates.ToDoList, error) {
	return c.toDoListRepo.FindAll()
}
