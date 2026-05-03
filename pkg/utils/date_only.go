package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type DateOnly struct {
	time.Time
}

const DateOnlyLayout = "2006-01-02"

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(DateOnlyLayout, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(DateOnlyLayout))), nil
}

func (d DateOnly) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time.Format(DateOnlyLayout), nil
}

func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		d.Time = v
	case string:
		t, err := time.Parse(DateOnlyLayout, v)
		if err != nil {
			return err
		}
		d.Time = t
	default:
		return fmt.Errorf("cannot scan type %T into DateOnly", value)
	}
	return nil
}
