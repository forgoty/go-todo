package main

import (
	"fmt"

	"github.com/forgoty/go-todo/pkg/todo/todolist/interfaces/rest"
)

func main() {
	fmt.Println(rest.NewTodoListController())
}
