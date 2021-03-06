package schema

import "errors"

var (
	ErrNotFound          = errors.New("schema: not found")
	ErrAlreadyExists     = errors.New("schema: already registered")
	ErrSecondaryKey      = errors.New("schema: secondary key not found in schema")
	ErrFieldTypeMismatch = errors.New("schema: field type mismatch")
	ErrBadDefinition     = errors.New("schema: invalid definition")
)
