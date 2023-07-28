package email

import (
	"errors"
	"regexp"
)

const (
	validatorRegex = `(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))`
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Email struct {
	value string
}

func New(str string) (Email, error) {
	if !valid(str) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: str}, nil
}

func (e Email) Value() string {
	return e.value
}

func valid(str string) bool {
	return regexp.MustCompile(validatorRegex).MatchString(str)
}
