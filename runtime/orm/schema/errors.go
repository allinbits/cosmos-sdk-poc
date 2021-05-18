package schema

import "errors"

var (
	ErrNotFound             = errors.New("schema: not found")
	ErrAlreadyExists        = errors.New("schema: already registered")
	ErrRegister             = errors.New("schema: error whilst registering new object schema")
	ErrSecondaryKey         = errors.New("schema: secondary key not found in schema")
	ErrFieldEncode          = errors.New("schema: unable to encode field")
	ErrEmptyFieldValue      = errors.New("schema: empty field value")
	ErrUnsupportedFieldKind = errors.New("schema: field kind is not supported for indexing")
	ErrFieldTypeMismatch    = errors.New("schema: field type mismatch")
)
