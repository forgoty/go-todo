package todolist

import (
	"github.com/forgoty/go-todo/pkg/todo/todolist/application/pkg/application"
	"github.com/forgoty/go-todo/pkg/todo/todolist/domain/models/aggregates"
)

type toDoListController struct {
	app        application.Application
	repository aggregates.ToDoListRepository
}
