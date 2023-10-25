package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type DueDate struct {
	time.Time
}

func (d DueDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	//*d = DueDate(t)
	d.Time = t
	return nil
}

func (d DueDate) MarshalJSON() ([]byte, error) {
	//return json.Marshal(time.Time(d))
	return json.Marshal(d.Time)
}

func (d DueDate) Scan(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		d.Time = v
		//*d = DueDate(v)
		return nil
	}
	return errors.New("Invalid DueDate value")
}

func (d DueDate) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}
