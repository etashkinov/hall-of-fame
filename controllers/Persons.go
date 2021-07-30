package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/persistence"
	"github.com/etashkinov/hall-of-fame/validator"
)

type PersonCreateRequest struct {
	Name string `json:"name" example:"Homer Simpson"`
}
type PersonUpdateRequest struct {
	Name        string `json:"name" example:"Homer Simpson"`
	Description string `json:"description" example:"Fat bold white guy"`
}

func CreatePerson(c *gin.Context) {
	var request PersonCreateRequest

	bindJSON(c, &request)

	if messagesError := validator.General(request, nil); messagesError != nil {
		badRequest(c, messagesError)
		return
	}

	person, err := persistence.CreatePerson(request.Name)

	response(c, person, err)
}

func UpdatePerson(c *gin.Context) {
	personId, err := getPersonId(c)
	if err != nil {
		serverError(c, err)
		return
	}

	var request PersonUpdateRequest
	bindJSON(c, &request)

	person, err := persistence.UpdatePerson(personId, request.Name, request.Description)

	response(c, person, err)
}

func GetPersons(c *gin.Context) {
	persons, err := persistence.GetPersons()
	response(c, persons, err)
}

func GetPerson(c *gin.Context) {
	personId, err := getPersonId(c)
	if err != nil {
		serverError(c, err)
		return
	}

	person, err := persistence.GetPerson(personId)
	response(c, person, err)
}

func getPersonId(c *gin.Context) (personId int64, err error) {
	param, err := strconv.Atoi(c.Param("id"))
	return int64(param), err
}

func response(c *gin.Context, body interface{}, err error) {
	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, body)
}
