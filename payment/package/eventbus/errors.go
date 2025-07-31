package eventbus

import "errors"

var (
	ErrHandlerAlreadyRegistered = errors.New("event handler already registered")
	ErrHandlerNotFound          = errors.New("event handler not found")
)
