package strtypes

import (
	"errors"
	"fmt"
)

var (
	ErrInvalid = errors.New("invalid value")
)

func InvalidTypeErr(expected string, actual interface{}) error {
	return fmt.Errorf("invalid type for %s: %T", expected, actual)
}
