package aggregates

import (
	"github.com/foroto/go-todo/internal/todo/todolist/domain/commands"
	"github.com/foroto/go-todo/internal/todo/todolist/domain/model/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewToDoList(t *testing.T) {
	command := commands.NewCreateToDoListCommand("test")
	todo := NewToDoList(command)
	assert.NotNil(t, todo)
}

func TestAddToDoItem(t *testing.T) {
	tdlist := ToDoList{
		Name: "test",
	}
	tditem := entities.ToDoItem{
		Name:        "test",
		Description: "test",
	}
	tdlist.AddToDoItem(tditem)
	assert.Equal(t, len(tdlist.ToDoItems), 1)
}

func TestAddToDoItems(t *testing.T) {
	tdlist := ToDoList{
		Name: "test",
	}
	tditem1 := entities.ToDoItem{
		Name:        "test",
		Description: "test",
	}
	tditem2 := entities.ToDoItem{
		Name:        "test",
		Description: "test",
	}
	tdlist.AddToDoItems([]entities.ToDoItem{tditem1, tditem2})
	assert.Equal(t, len(tdlist.ToDoItems), 2)
}
