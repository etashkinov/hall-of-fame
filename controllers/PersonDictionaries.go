package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/persistence"
)

type dictRequest struct {
	DictId      int64
	Since       persistence.Date
	Till        persistence.Date
	Level       persistence.Level
	TeamId      int64
	Description string
}

func AddPersonDict(c *gin.Context) {
	personId, _ := getId(c)
	dictType := getDictType(c)

	skill, err := add(c, personId, dictType)

	ok(c, skill, err)
}

func add(c *gin.Context, personId int64, dictType string) (dict interface{}, err error) {
	var request dictRequest
	bindJSON(c, &request)

	switch dictType {
	case "skills":
		return persistence.AddPersonSkill(personId, request.DictId, request.Since, request.Level)
	case "achievements":
		return persistence.AddPersonAchievement(personId, request.DictId, request.Since, request.Description)
	case "expertises":
		return persistence.AddPersonExpertise(personId, request.DictId, request.Since, request.Level)
	case "positions":
		return persistence.AddPersonPosition(personId, request.DictId, request.TeamId, request.Since, request.Description)
	}

	return nil, unknownDictType(dictType)
}

func GetPersonDicts(c *gin.Context) {
	personId, _ := getId(c)
	dictType := getDictType(c)

	dicts, err := get(personId, dictType)

	ok(c, dicts, err)
}

func get(personId int64, dictType string) (interface{}, error) {
	switch dictType {
	case "skills":
		return persistence.GetPersonSkills(personId)
	case "achievements":
		return persistence.GetPersonAchievements(personId)
	case "expertises":
		return persistence.GetPersonExpertises(personId)
	case "positions":
		return persistence.GetPersonPositions(personId)
	}

	return nil, unknownDictType(dictType)
}

func unknownDictType(dictType string) error {
	return fmt.Errorf("unknown dict type  %s", dictType)
}

func UpdatePersonDict(c *gin.Context) {
	personDictId, _ := getIntParam(c, "dictId")
	dictType := getDictType(c)
	skill, err := update(c, personDictId, dictType)

	ok(c, skill, err)
}

func update(c *gin.Context, personDictId int64, dictType string) (dict interface{}, err error) {
	var request dictRequest
	bindJSON(c, &request)

	switch dictType {
	case "skills":
		return persistence.UpdatePersonSkill(personDictId, request.Since, request.Level)
	case "achievements":
		return persistence.UpdatePersonAchievement(personDictId, request.Since, request.Description)
	case "expertises":
		return persistence.UpdatePersonExpertise(personDictId, request.Since, request.Level)
	case "positions":
		return persistence.UpdatePersonPosition(personDictId, request.Since, request.Till, request.Description)
	}

	return nil, unknownDictType(dictType)
}

func DeletePersonDict(c *gin.Context) {
	personDictId, _ := getIntParam(c, "dictId")
	dictType := getDictType(c)
	err := delete(personDictId, dictType)

	ok(c, nil, err)
}

func delete(personDictId int64, dictType string) error {
	switch dictType {
	case "skills":
		return persistence.DeletePersonSkill(personDictId)
	case "achievements":
		return persistence.DeletePersonAchievement(personDictId)
	case "expertises":
		return persistence.DeletePersonExpertise(personDictId)
	case "positions":
		return persistence.DeletePersonPosition(personDictId)
	}

	return unknownDictType(dictType)
}
