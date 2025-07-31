package commandbus

import "errors"

var (
	ErrHandlerAlreadyRegistered = errors.New("command handler already registered")
	ErrHandlerNotFound          = errors.New("command handler not found")
)
