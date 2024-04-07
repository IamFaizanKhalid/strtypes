package time

import (
	"database/sql/driver"
	"fmt"
	"github.com/IamFaizanKhalid/strtypes"
	"strings"
	"time"
)

type Time string

const timeFormat = time.RFC3339

func (u Time) String() string {
	return string(u)
}

func (u Time) Valid() bool {
	_, err := time.Parse(timeFormat, string(u))
	return err == nil
}

func (u Time) BuiltIn() time.Time {
	t, _ := time.Parse(timeFormat, string(u))
	return t
}

func (u Time) Date() Date {
	return Date(u.BuiltIn().Format(dateFormat))
}

func (u Time) Year() int {
	return u.BuiltIn().Year()
}

func (u Time) Month() time.Month {
	return u.BuiltIn().Month()
}

func (u Time) Day() int {
	return u.BuiltIn().Day()
}

func (u Time) Weekday() time.Weekday {
	return u.BuiltIn().Weekday()
}

func (u Time) Hour() int {
	return u.BuiltIn().Hour()
}

func (u Time) Minute() int {
	return u.BuiltIn().Minute()
}

func (u Time) Second() int {
	return u.BuiltIn().Second()
}

func (u Time) Nanosecond() int {
	return u.BuiltIn().Nanosecond()
}

func (u Time) YearDay() int {
	return u.BuiltIn().YearDay()
}

/////// Constructors ///////

func Now() Time {
	return Time(time.Now().Format(timeFormat))
}

func TimeFromBuiltIn(t time.Time) Time {
	return Time(t.Format(timeFormat))
}

/////// SQL ///////

func (u Time) Value() (driver.Value, error) {
	if u == "" {
		return nil, nil
	}

	if !u.Valid() {
		return nil, strtypes.ErrInvalid
	}
	return string(u), nil
}

func (u *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s, ok := value.(time.Time)
	if !ok {
		return strtypes.InvalidTypeErr("Time", value)
	}

	tm := TimeFromBuiltIn(s)
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
