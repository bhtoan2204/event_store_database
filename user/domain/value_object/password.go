package value_object

import (
	"errors"
	"regexp"
)

type Password struct {
	value string
}

func NewAndValidatePassword(value string) (*Password, error) {
	if len(value) < 8 || len(value) > 20 {
		return nil, errors.New("password must be between 8 and 20 characters")
	}

	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#~$%^&*()_+{}\[\]:;<>,.?\/\\|-]`).MatchString
	)

	if !hasUpper(value) {
		return nil, errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower(value) {
		return nil, errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber(value) {
		return nil, errors.New("password must contain at least one digit")
	}
	if !hasSpecial(value) {
		return nil, errors.New("password must contain at least one special character")
	}

	return &Password{value: value}, nil
}

func (p *Password) Value() string {
	return p.value
}
