package time

import (
	"database/sql/driver"
	"fmt"
	"github.com/IamFaizanKhalid/strtypes"
	"strings"
	"time"
)

type Time string

func (u Time) String() string {
	return string(u)
}

func (u Time) Valid() bool {
	_, err := time.Parse(time.RFC3339, string(u))
	return err == nil
}

/////// Constructors ///////

func NewTime() Time {
	return Time(time.Now().Format(time.RFC3339))
}

func NewTimePointer() *Time {
	tm := Time(time.Now().Format(time.RFC3339))
	return &tm
}

/////// SQL ///////

func (u Time) Value() (driver.Value, error) {
	if !u.Valid() {
		return nil, strtypes.ErrInvalid
	}
	return string(u), nil
}

func (u *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s, ok := value.(string)
	if !ok {
		return strtypes.InvalidTypeErr("Time", s)
	}

	tm := Time(s)
	if !tm.Valid() {
		return strtypes.ErrInvalid
	}

	*u = tm

	return nil
}

/////// JSON ///////

func (u Time) MarshalJSON() ([]byte, error) {
	if !u.Valid() {
		return nil, strtypes.ErrInvalid
	}
	return []byte(fmt.Sprintf("\"%s\"", u)), nil
}

func (u *Time) UnmarshalJSON(data []byte) error {
	var s = strings.Trim(string(data), `"`)

	tm := Time(s)
	if !tm.Valid() {
		return strtypes.ErrInvalid
	}

	*u = tm

	return nil
}
