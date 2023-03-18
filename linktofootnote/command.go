package linktofootnote

import (
	"errors"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/io"
)

var (
	SourceFileDoesNotExistErr       = errors.New("source file does not exist")
	DestinationFileAlreadyExistsErr = errors.New("destination file already exists")
)

type Command struct {
	fileSystem *io.FileSystem
}

func (c Command) Run(sourceFilepath, destinationFilepath string) error {
	if !c.fileSystem.FileExists(sourceFilepath) {
		return SourceFileDoesNotExistErr
	}
	if c.fileSystem.FileExists(destinationFilepath) {
		return DestinationFileAlreadyExistsErr
	}
	_ = c.fileSystem.ReadFile(sourceFilepath)
	panic("not implemented")
}

func NewCommand(fileSystem *io.FileSystem) *Command {
	return &Command{fileSystem: fileSystem}
}
