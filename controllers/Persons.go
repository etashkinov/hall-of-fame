package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/persistence"
	"github.com/etashkinov/hall-of-fame/validator"
)

type personRequest struct {
	Name        string `json:"name" example:"Homer Simpson"`
	Description string `json:"description" example:"Fat bold white guy"`
}

func CreatePerson(c *gin.Context) {
	var request personRequest

	bindJSON(c, &request)

	if messagesError := validator.General(request, nil); messagesError != nil {
		badRequest(c, messagesError)
		return
	}

	person, err := persistence.CreatePerson(request.Name, request.Description)

	ok(c, person, err)
}

func UpdatePerson(c *gin.Context) {
	personId, err := getId(c)
	if err != nil {
		serverError(c, err)
		return
	}

	var request personRequest
	bindJSON(c, &request)

	person, err := persistence.UpdatePerson(personId, request.Name, request.Description)

	ok(c, person, err)
}

func GetPersons(c *gin.Context) {
	persons, err := persistence.GetPersons()
	ok(c, persons, err)
}

func GetPerson(c *gin.Context) {
	personId, err := getId(c)
	if err != nil {
		serverError(c, err)
		return
	}

	person, err := persistence.GetPerson(personId)
	ok(c, person, err)
}
