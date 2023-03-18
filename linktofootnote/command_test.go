package linktofootnote_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/io"
	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/linktofootnote"
	"github.com/stretchr/testify/assert"
)

func TestCommand_Run_ShouldCreateFileWithTransformedText(t *testing.T) {
	command := linktofootnote.NewCommand(io.NewFileSystem())
	err := command.Run("../testdata/test.md", "../testdata/destination.md")
	assert.NoError(t, err)
}

func TestCommand_Run_ShouldReturnError_WhenSourceFileDoesNotExist(t *testing.T) {
	command := linktofootnote.NewCommand(io.NewFileSystem())
	err := command.Run("../testdata/noexistingfile.md", "../testdata/destination.md")
	assert.EqualError(t, err, "source file does not exist")
}

func TestCommand_Run_ShouldReturnError_WhenDestinationFileAlreadyExists(t *testing.T) {
	command := linktofootnote.NewCommand(io.NewFileSystem())
	err := command.Run("../testdata/test.md", "../testdata/test.md")
	assert.EqualError(t, err, "destination file already exists")
}
