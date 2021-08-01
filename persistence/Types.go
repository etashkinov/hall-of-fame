package persistence

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Level int

const (
	LEVEL_EXPLORER Level = 0
	LEVEL_HABITANT Level = 1
	LEVEL_GURU     Level = 2
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

type Dictionary struct {
	Id          int64  `json:"dictId" db:"dict_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PersonSkill struct {
	Id        int64 `json:"id"`
	PersionId int64 `json:"personId" db:"person_id"`
	Since     Date  `json:"since"`
	Level     Level `json:"level"`
	Dictionary
}

type PersonAchievement struct {
	Id          int64
	PersionId   int64
	Since       Date
	Description string
	Dictionary
}

type PersonExpertise struct {
	Id        int64
	PersionId int64
	Since     string
	Level     Level
	Dictionary
}

type PersonPosition struct {
	Id          int64
	PersionId   int64
	TeamId      string
	Since       Date
	Till        Date
	Description string
	Dictionary
}

type Person struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
