package domain

import "errors"

var (
    ErrFileNotFound = errors.New("file not found")
    ErrInvalidPath  = errors.New("invalid file path")
)