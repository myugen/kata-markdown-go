package linktofootnote

type Command struct{}

func (c Command) Run(sourceFilepath, destinationFilepath string) error {
	panic("not implemented yet")
}

func NewCommand() *Command {
	return &Command{}
}
