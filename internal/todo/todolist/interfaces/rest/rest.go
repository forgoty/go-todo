package rest

import (
	"github.com/forgoty/go-todo/internal/todo/todolist/application"
	"github.com/forgoty/go-todo/internal/todo/todolist/domain/model/aggregates"
)

type toDoListController struct {
	app        *application.Application
	repository *aggregates.ToDoListRepository
}

func NewTodoListController() *toDoListController {
	return &toDoListController{
		app:        nil,
		repository: nil,
	}
}
