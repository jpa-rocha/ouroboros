// Package utils holds utility function that may be useful to other packages
package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

const (
	OwnerReadWriteExecute = uint32(493)
)

type ContextKey string

func (c *ContextKey) String() string {
	return string(*c)
}

func WriteToFile(path string, content []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := file.Write(content); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func CreateFolder(path string) error {
	err := os.Mkdir(path, fs.FileMode(OwnerReadWriteExecute))
	if err != nil {
		if noPermission := errors.Is(err, fs.ErrPermission); noPermission {
			return ErrPathNotAllowed
		} else if folderExists := errors.Is(err, fs.ErrExist); folderExists {
			return ErrFolderAlreadyExists
		}

		return fmt.Errorf("%w", err)
	}

	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func RemoveFolder(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return ErrCouldNotDelete
	}

	return nil
}
