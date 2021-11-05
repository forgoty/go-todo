package aggregates

import (
	"github.com/forgoty/go-todo/pkg/todo/todolist/domain/commands"
	"github.com/forgoty/go-todo/pkg/todo/todolist/domain/model/entities"
	"github.com/google/uuid"
)

type ToDoList struct {
	Id        uuid.UUID
	Name      string
	ToDoItems []entities.ToDoItem
}

func NewToDoList(command *commands.CreateToDoListCommand) *ToDoList {
	return &ToDoList{
		Id:   uuid.New(),
		Name: command.Name,
	}
}

func (tdl *ToDoList) AddToDoItem(todo entities.ToDoItem) {
	tdl.ToDoItems = append(tdl.ToDoItems, todo)
}

func (tdl *ToDoList) AddToDoItems(todos []entities.ToDoItem) {
	tdl.ToDoItems = append(tdl.ToDoItems, todos...)
}
