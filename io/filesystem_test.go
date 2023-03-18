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
