package persistence

import (
	"time"
)

type Level string

const (
	LEVEL_EXPLORER Level = "EXPLORER"
	LEVEL_HABITANT Level = "HABITANT"
	LEVEL_GURU     Level = "GURU"
)

type Dictionary struct {
	Id          int64
	Name        string
	Description string
}

type PersonSkill struct {
	Id        int64
	PersionId int64
	Skill     Dictionary
	Since     time.Time
	Level     Level
}

type PersonAchievement struct {
	Id          int64
	PersionId   int64
	Achievement Dictionary
	Since       time.Time
	Description string
}

type PersonExpertise struct {
	Id        int64
	PersionId int64
	Expertise Dictionary
	Since     time.Time
	Level     Level
}

type PersonPosition struct {
	Id          int64
	PersionId   int64
	Position    Dictionary
	TeamId      string
	Since       time.Time
	Till        time.Time
	Description string
}

type Person struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
