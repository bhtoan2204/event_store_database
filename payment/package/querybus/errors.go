package querybus

import "errors"

var (
	ErrHandlerAlreadyRegistered = errors.New("query handler already registered")
	ErrHandlerNotFound          = errors.New("query handler not found")
)
