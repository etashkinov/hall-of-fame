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

type Skill struct {
	Id          int64
	Name        string
	Description string
}

type PersonSkill struct {
	Id        int64
	PersionId int64
	Skill     Skill
	Since     time.Time
	Level     Level
}

type Achievement struct {
	Id          int64
	Name        string
	Description string
}

type PersonAchievement struct {
	Id          int64
	PersionId   int64
	Achievement Achievement
	Since       time.Time
	Description string
}

type Expertise struct {
	Id          int64
	Name        string
	Description string
}

type PersonExpertise struct {
	Id        int64
	PersionId int64
	Expertise Expertise
	Since     time.Time
	Level     Level
}

type Position struct {
	Id          int64
	Name        string
	Description string
}

type PersonPosition struct {
	Id          int64
	PersionId   int64
	Position    Position
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
