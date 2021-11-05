package application

import (
	"fmt"
	"github.com/forgoty/go-todo/pkg/todo/todolist/application/pkg/commandservices"
	"github.com/forgoty/go-todo/pkg/todo/todolist/application/pkg/queryservices"
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
