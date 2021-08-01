package persistence

import (
	"time"
)

type Level int

const (
	LEVEL_EXPLORER Level = 0
	LEVEL_HABITANT Level = 1
	LEVEL_GURU     Level = 2
)

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
	Id        int64  `json:"id"`
	PersionId int64  `json:"personId" db:"person_id"`
	Since     Date   `json:"since"`
	Details   string `json:"details" db:"description"`
	Dictionary
}

type PersonExpertise struct {
	Id        int64 `json:"id"`
	PersionId int64 `json:"personId" db:"person_id"`
	Since     Date  `json:"since"`
	Level     Level `json:"level"`
	Dictionary
}

type PersonPosition struct {
	Id        int64  `json:"id"`
	PersionId int64  `json:"personId" db:"person_id"`
	TeamId    int64  `json:"teamId"  db:"team_id"`
	Since     Date   `json:"since"`
	Till      Date   `json:"till"`
	Details   string `json:"description" db:"description"`
	Dictionary
}

type Person struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
