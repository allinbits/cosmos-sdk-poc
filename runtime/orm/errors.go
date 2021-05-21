package orm

import "errors"

var (
	ErrNotFound = errors.New("orm: object not found")
)
