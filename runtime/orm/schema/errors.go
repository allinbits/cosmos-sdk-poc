package schema

import "errors"

var (
	ErrNotFound      = errors.New("schema: not found")
	ErrAlreadyExists = errors.New("schema: already registered")
	ErrRegister      = errors.New("schema: error whilst registering new object schema")
)
