package linktofootnote_test

import (
	"os"
	"testing"

	"github.com/myugen/kata-markdown-go/io"
	"github.com/myugen/kata-markdown-go/linktofootnote"
	"github.com/stretchr/testify/assert"
)

func TestCommand_Run_ShouldCreateFileWithTransformedText(t *testing.T) {
	command := linktofootnote.NewCommand(io.NewFileSystem())
	source := "../testdata/test.md"
	destination := "../testdata/destination.md"
	err := command.Run(source, destination)
	if assert.NoError(t, err) {
		content, _ := os.ReadFile(destination)
		assert.Equal(t, "# Test file\n\nEl libro de Código Sostenible[^anchor1] es un librazo.\n¡Cómpralo!\n\n[^anchor1]: https://www.codigosostenible.com", string(content))
	}

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
