package carplate

import (
	"errors"
	"regexp"
	"strings"
)

var (
	InvalidCarPlate = errors.New("invalid car plate")
)

type CarPlate struct {
	value string
}

func New(str string) (CarPlate, error) {
	str = strings.ToUpper(str)
	if !valid(str) {
		return CarPlate{}, InvalidCarPlate
	}
	return CarPlate{str}, nil
}

func (c CarPlate) Value() string {
	return c.value
}

func valid(str string) bool {
	return regexp.MustCompile("[a-zA-Z]{3}\\d{4}$").MatchString(str)
}
