package linktofootnote

import (
	"errors"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/io"
	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/markdown"
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
	transformer := markdown.NewTransformer(c.fileSystem.ReadFile(sourceFilepath))
	transformedText := transformer.Transform()
	return c.fileSystem.WriteFile(destinationFilepath, transformedText)
}

func NewCommand(fileSystem *io.FileSystem) *Command {
	return &Command{fileSystem: fileSystem}
}
