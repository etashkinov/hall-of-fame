package persistence

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (date Date) Format() string {
	return string(fmt.Sprintf("%04d-%02d-%02d", date.Year, date.Month, date.Day))
}

func (date *Date) Parse(value string) (err error) {
	split := strings.Split(value, "-")
	year, err := strconv.Atoi(split[0])
	if err != nil {
		return
	}

	month, err := strconv.Atoi(split[1])
	if err != nil {
		return
	}

	day, err := strconv.Atoi(split[2])
	if err != nil {
		return
	}

	*date = Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
	return nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(date.Format())
}

func (date *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return date.Parse(s)
}

func (date Date) Value() (driver.Value, error) {
	return date.Format(), nil
}

func (date *Date) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil {
		return nil
	}
	if str, err := driver.String.ConvertValue(value); err == nil {
		// if this is a bool type
		if v, ok := str.(string); ok {
			dayPart := strings.Fields(v)[0]
			err = date.Parse(dayPart)
			return err
		}
	}
	// otherwise, return an error
	return fmt.Errorf("failed to scan %s date", value)
}
