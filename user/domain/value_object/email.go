package value_object

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(value) {
		return nil, errors.New("invalid email")
	}
	return &Email{value: value}, nil
}

func (e *Email) Value() string {
	return e.value
}
