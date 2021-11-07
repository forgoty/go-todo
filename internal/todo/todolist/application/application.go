package application

import (
	"github.com/forgoty/go-todo/internal/todo/todolist/application/internal/commandservices"
	"github.com/forgoty/go-todo/internal/todo/todolist/application/internal/queryservices"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateToDoList commandservices.CreateToDoListCommandHandler
	DeleteToDoList commandservices.CreateToDoListCommandHandler
}

type Queries struct {
	FindAllToDoLists queryservices.FindAllToDoListByIdQueryHandler
	FindToDoLists    queryservices.FindToDoListByIdQueryHandler
}
