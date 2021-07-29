package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/models"
	"github.com/etashkinov/hall-of-fame/validator"
)

type CreatePersonRequest struct {
	Name string `json:"name" example:"Homer Simpson"`
}

func CreatePerson(c *gin.Context) {
	var request CreatePersonRequest

	bindJSON(c, &request)

	if messagesError := validator.General(request, nil); messagesError != nil {
		badRequest(c, messagesError)
		return
	}

	person, err := models.CreatePerson(request.Name)

	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, person)
}

func GetPerson(c *gin.Context) {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		serverError(c, err)
		return
	}

	person, err := models.GetPerson(int64(personId))
	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, person)
}
