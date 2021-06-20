package commands

type CreateToDoListCommand struct {
	Name string
}

func NewCreateToDoListCommand(name string) *CreateToDoListCommand {
	return &CreateToDoListCommand{
		Name: name,
	}
}
