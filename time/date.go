package time

import (
	"database/sql/driver"
	"fmt"
	"github.com/IamFaizanKhalid/strtypes"
	"strings"
	"time"
)

type Date string

const dateFormat = time.DateOnly

func (u Date) String() string {
	return string(u)
}

func (u Date) Valid() bool {
	_, err := time.Parse(dateFormat, string(u))
	return err == nil
}

/////// Constructors ///////

func Today() Date {
	return Date(time.Now().Format(dateFormat))
}

func TodayPointer() *Date {
	tm := Date(time.Now().Format(dateFormat))
	return &tm
}

/////// SQL ///////

func (u Date) Value() (driver.Value, error) {
	if !u.Valid() {
		return nil, strtypes.ErrInvalid
	}
	return string(u), nil
}

func (u *Date) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s, ok := value.(string)
	if !ok {
		return strtypes.InvalidTypeErr("Date", s)
	}

	tm := Date(s)
	if !tm.Valid() {
		return strtypes.ErrInvalid
	}

	*u = tm

	return nil
}

/////// JSON ///////

func (u Date) MarshalJSON() ([]byte, error) {
	if !u.Valid() {
		return nil, strtypes.ErrInvalid
	}
	return []byte(fmt.Sprintf("\"%s\"", u)), nil
}

func (u *Date) UnmarshalJSON(data []byte) error {
	var s = strings.Trim(string(data), `"`)

	tm := Date(s)
	if !tm.Valid() {
		return strtypes.ErrInvalid
	}

	*u = tm

	return nil
}