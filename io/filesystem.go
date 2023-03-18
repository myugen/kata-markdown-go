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

func (s FileSystem) WriteFile(filepath, text string) error {
	return os.WriteFile(filepath, []byte(text), 0666)
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}
