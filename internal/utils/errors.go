package utils

import "errors"

var (
	ErrPathNotAllowed      = errors.New("error: path not allowed")
	ErrFolderAlreadyExists = errors.New("error: folder already exists")
	ErrCouldNotDelete      = errors.New("error: could not delete folder")
	ErrCreatingTemplate    = errors.New("error: creating template")
)
