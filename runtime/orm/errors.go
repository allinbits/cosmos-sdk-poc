package orm

import "errors"

var (
	ErrNotFound      = errors.New("orm: object not found")
	ErrAlreadyExists = errors.New("orm: object already exists")
)
