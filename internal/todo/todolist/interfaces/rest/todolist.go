package todolist

import (
	"github.com/foroto/go-todo/internal/todo/todolist/application/internal/application"
	"github.com/foroto/go-todo/internal/todo/todolist/domain/models/aggregates"
)

type toDoListController struct {
	app        application.Application
	repository aggregates.ToDoListRepository
}
