package errs

import "fmt"

var (
	ErrNotFound = fmt.Errorf("resource not found")
)
