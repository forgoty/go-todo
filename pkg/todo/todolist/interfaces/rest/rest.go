package rest

import (
	"github.com/forgoty/go-todo/pkg/todo/todolist/application"
	"github.com/forgoty/go-todo/pkg/todo/todolist/domain/model/aggregates"
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
