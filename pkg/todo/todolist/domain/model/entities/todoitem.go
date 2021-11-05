package entities

type ToDoItem struct {
	Name        string
	Description string
}

func NewToDoItem() *ToDoItem {
	return &ToDoItem{
		Name:        "test",
		Description: "test",
	}
}
