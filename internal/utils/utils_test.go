package utils_test

import (
	"os"
	"testing"

	utils "ouroboros/internal/utils"

	"github.com/matryer/is"
)

func TestWriteToFile(t *testing.T) {
	tables := []struct {
		err      error
		name     string
		filePath string
		content  string
	}{
		{
			name:     "CheckWrite",
			filePath: "./test.txt",
			content:  "Hi, im the test content.",
			err:      nil,
		},
		{
			name:     "CheckError",
			filePath: "./test.txt",
			content:  "Hi, im the test content.",
			err:      os.ErrExist,
		},
	}

	is := is.New(t)

	for _, tt := range tables {
		t.Run(tt.name, func(_ *testing.T) {
			err := utils.WriteToFile(tt.filePath, []byte(tt.content))
			if err != nil {
				is.Equal(err, tt.err)
			} else {
				result, _ := os.ReadFile(tt.filePath)
				_ = os.Remove(tt.filePath)
				is.Equal(string(result), tt.content)
				is.NoErr(err)
			}
		})
	}
}

func TestCreateFolder(t *testing.T) {
	tables := []struct {
		err        error
		name       string
		folderPath string
	}{
		{
			name:       "CreateFolder",
			folderPath: "./test",
			err:        nil,
		},
		{
			name:       "CheckError_00",
			folderPath: "./test",
			err:        utils.ErrFolderAlreadyExists,
		},
		{
			name:       "CheckError_01",
			folderPath: "/root/test",
			err:        utils.ErrPathNotAllowed,
		},
	}

	is := is.New(t)

	for _, tt := range tables {
		t.Run(tt.name, func(_ *testing.T) {
			err := utils.CreateFolder(tt.folderPath)
			if err != nil {
				_ = os.Remove(tt.folderPath)

				is.Equal(err, tt.err)
			} else {
				is.NoErr(err)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	tables := []struct {
		expect   bool
		name     string
		filePath string
	}{
		{
			name:     "CheckExists",
			filePath: "./test.txt",
			expect:   true,
		},
		{
			name:     "CheckDoesntExist",
			filePath: "./test.txt",
			expect:   false,
		},
	}

	is := is.New(t)

	_ = utils.WriteToFile(tables[0].filePath, []byte("test"))
	for _, tt := range tables {
		t.Run(tt.name, func(_ *testing.T) {
			ok := utils.FileExists(tt.filePath)
			is.Equal(true, ok)
		})
	}
	_ = os.Remove(tables[0].filePath)
}
