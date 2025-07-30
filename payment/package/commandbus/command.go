package commandbus

type ICommand interface {
	CommandName() string
}

type command struct {
	name string
}

func NewCommand(name string) ICommand {
	return &command{
		name: name,
	}
}

func (c *command) CommandName() string {
	return c.name
}
