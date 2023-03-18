package main

import (
	"errors"
	"log"
	"os"

	"github.com/myugen/kata-markdown-go/io"
	"github.com/myugen/kata-markdown-go/linktofootnote"
)

var (
	notEnoughArgumentsErr = errors.New("not enough arguments provided, 2 expected")
)

func main() {
	argsWithoutProgram, err := parseArguments()
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	sourceFilepath := argsWithoutProgram[0]
	destinationFilepath := argsWithoutProgram[1]

	fileSystem := io.NewFileSystem()
	command := linktofootnote.NewCommand(fileSystem)

	if err = command.Run(sourceFilepath, destinationFilepath); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func parseArguments() ([]string, error) {
	argsWithoutProgram := os.Args[1:]
	if len(argsWithoutProgram) < 2 {
		return nil, notEnoughArgumentsErr
	}
	return argsWithoutProgram, nil
}
