package io

import "os"

type FileSystem struct{}

func (FileSystem) FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}
