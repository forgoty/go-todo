package commandservices

import (
	"github.com/forgoty/go-todo/pkg/todo/todolist/domain/commands"
	"github.com/forgoty/go-todo/pkg/todo/todolist/domain/model/aggregates"
)

type CreateToDoListCommandHandler struct {
	toDoListRepo aggregates.ToDoListRepository
}

func NewCreateToDoListCommandHandler(repo aggregates.ToDoListRepository) CreateToDoListCommandHandler {
	return CreateToDoListCommandHandler{
		toDoListRepo: repo,
	}
}

func (c *CreateToDoListCommandHandler) Handle(command *commands.CreateToDoListCommand) error {
	todolist := aggregates.NewToDoList(command)
	return c.toDoListRepo.Store(*todolist)
}
