package io

import "os"

type FileSystem struct{}

func (FileSystem) ReadFile(filepath string) string {
	content, _ := os.ReadFile(filepath)
	return string(content)

}

func (FileSystem) FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}
