package runtime

import "errors"

var (
	// BadRequest is returned when the request is malformed
	BadRequest = errors.New("runtime: bad request")
	// NotFound is returned when an object is not found
	NotFound = errors.New("runtime: not found")
)

func IsNotFound(err error) bool {
	return errors.Is(err, NotFound)
}
