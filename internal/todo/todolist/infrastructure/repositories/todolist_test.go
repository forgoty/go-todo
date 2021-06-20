package todolist

import (
	"github.com/foroto/go-todo/internal/todo/todolist/domain/commands"
	"github.com/foroto/go-todo/internal/todo/todolist/domain/model/aggregates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToDoList_FindById(t *testing.T) {
	repo := NewInMemoryToDoListRepository()
	command := commands.NewCreateToDoListCommand("test")
	todo := aggregates.NewToDoList(command)
	_ = repo.Store(*todo)
	savedTodo, err := repo.FindById(todo.Id)
	assert.Nil(t, err)
	assert.Equal(t, savedTodo, todo)
}

func TestToDoList_DeleteById(t *testing.T) {
	repo := NewInMemoryToDoListRepository()
	command := commands.NewCreateToDoListCommand("test")
	todo := aggregates.NewToDoList(command)
	_ = repo.Store(*todo)
	err := repo.DeleteById(todo.Id)
	assert.Nil(t, err)
	_, err = repo.FindById(todo.Id)
	assert.NotNil(t, err)
}

func TestToDoList_GetAll(t *testing.T) {
	repo := NewInMemoryToDoListRepository()
	command := commands.NewCreateToDoListCommand("test")
	todo1 := aggregates.NewToDoList(command)
	_ = repo.Store(*todo1)
	todo2 := aggregates.NewToDoList(command)
	_ = repo.Store(*todo2)
	slice, err := repo.FindAll()
	assert.Nil(t, err)
	assert.Len(t, slice, 2)
}
