package uuid

import (
	"database/sql/driver"
	"fmt"
	"github.com/IamFaizanKhalid/strtypes"
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

func New() UUID {
	return UUID(uuid.New().String())
}

func NewPointer() *UUID {
	uid := UUID(uuid.New().String())
	return &uid
}

/////// SQL ///////

func (u UUID) Value() (driver.Value, error) {
	if !u.Valid() {
		return nil, strtypes.ErrInvalid
	}
	return string(u), nil
}

func (u *UUID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s, ok := value.(string)
	if !ok {
		return strtypes.InvalidTypeErr("UUID", s)
	}

	uid := UUID(s)
	if !uid.Valid() {
		return strtypes.ErrInvalid
	}

	*u = uid

	return nil
}

/////// JSON ///////

func (u UUID) MarshalJSON() ([]byte, error) {
	if !u.Valid() {
		return nil, strtypes.ErrInvalid
	}
	return []byte(fmt.Sprintf("\"%s\"", u)), nil
}

func (u *UUID) UnmarshalJSON(data []byte) error {
	var s = strings.Trim(string(data), `"`)

	uid := UUID(s)
	if !uid.Valid() {
		return strtypes.ErrInvalid
	}

	*u = uid

	return nil
}
