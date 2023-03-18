package io_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/io"
	"github.com/stretchr/testify/assert"
)

func TestFileSystem_FileExists(t *testing.T) {
	type args struct {
		filepath string
	}
	testcases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false when file is not present in provided path",
			args: args{"../testdata/foo.md"},
			want: false,
		},
		{
			name: "should return true when file is present in provided path",
			args: args{"../testdata/test.md"},
			want: true,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			fileSystem := io.NewFileSystem()
			got := fileSystem.FileExists(testcase.args.filepath)
			assert.Equal(t, got, testcase.want)
		})
	}
}

func TestFileSystem_ReadFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return the content of provided file",
			args: args{"../testdata/test.md"},
			want: "# Test file\n\nEl [libro de Código Sostenible](www.codigosostenible.com) es un librazo.\n¡Cómpralo!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileSystem := io.NewFileSystem()
			got := fileSystem.ReadFile(tt.args.filepath)
			assert.Equal(t, tt.want, got)
		})
	}
}
