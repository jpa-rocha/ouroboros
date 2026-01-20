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

// ContextKey is a custom type for context keys used to pass values in context.
type ContextKey string

// String returns the string representation of the ContextKey.
func (c *ContextKey) String() string {
	return string(*c)
}

// WriteToFile creates a file at the given path and writes content to it.
// Takes file path and byte content to write.
// Returns error if file creation or write operation fails.
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

// CreateFolder creates a directory at the given path with owner read/write/execute permissions.
// Takes directory path to create.
// Returns ErrPathNotAllowed if permission is denied or ErrFolderAlreadyExists if directory exists.
// Returns error for other OS errors.
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

// FileExists checks if a file or directory exists at the given path.
// Takes file path to check.
// Returns true if file exists, false otherwise.
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

// RemoveFolder recursively deletes a directory and all its contents.
// Takes directory path to remove.
// Returns ErrCouldNotDelete if removal fails, nil on success.
func RemoveFolder(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return ErrCouldNotDelete
	}

	return nil
}
