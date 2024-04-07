package strtypes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type UUID string

func (u UUID) String() string {
	return string(u)
}

func (u UUID) Valid() bool {
	_, err := uuid.Parse(string(u))
	return err == nil
}

/////// Constructors ///////

func NewUUID() UUID {
	return UUID(uuid.New().String())
}

func NewUUIDPointer() *UUID {
	uid := UUID(uuid.New().String())
	return &uid
}

/////// SQL ///////

func (u UUID) Value() (driver.Value, error) {
	if !u.Valid() {
		return nil, ErrInvalidUUID
	}
	return string(u), nil
}

func (u *UUID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s, ok := value.(string)
	if !ok {
		return InvalidTypeErr(s)
	}

	uid := UUID(s)
	if !uid.Valid() {
		return ErrInvalidUUID
	}

	*u = uid

	return nil
}

/////// JSON ///////

func (u UUID) MarshalJSON() ([]byte, error) {
	if !u.Valid() {
		return nil, ErrInvalidUUID
	}
	return []byte(fmt.Sprintf("\"%s\"", u)), nil
}

func (u *UUID) UnmarshalJSON(data []byte) error {
	var s = strings.Trim(string(data), `"`)

	uid := UUID(s)
	if !uid.Valid() {
		return ErrInvalidUUID
	}

	*u = uid

	return nil
}

/////// Errors ///////

var (
	ErrInvalidUUID = errors.New("invalid UUID")
)

func InvalidTypeErr(s string) error {
	return fmt.Errorf("invalid type for UUID: %T", s)
}
