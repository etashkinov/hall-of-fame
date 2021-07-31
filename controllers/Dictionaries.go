package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/persistence"
	"github.com/etashkinov/hall-of-fame/validator"
)

type dictionaryRequest struct {
	Name        string
	Description string
}

func CreateDictionary(c *gin.Context) {
	var request dictionaryRequest

	bindJSON(c, &request)

	if messagesError := validator.General(request, nil); messagesError != nil {
		badRequest(c, messagesError)
		return
	}

	dictType := getDictType(c)
	dict, err := persistence.CreateDictionary(dictType, request.Name, request.Description)

	ok(c, dict, err)
}

func UpdateDictionary(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		serverError(c, err)
		return
	}

	var request dictionaryRequest
	bindJSON(c, &request)

	dictType := getDictType(c)
	person, err := persistence.UpdateDictionary(dictType, id, request.Name, request.Description)

	ok(c, person, err)
}

func GetDictionaries(c *gin.Context) {
	dictType := getDictType(c)
	dicts, err := persistence.GetDictionaries(dictType)
	ok(c, dicts, err)
}

func DeleteDictionary(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		serverError(c, err)
		return
	}

	dictType := c.Param("type")
	dict, err := persistence.DeleteDictionary(dictType, id)
	ok(c, dict, err)
}

func getDictType(c *gin.Context) string {
	return c.Param("type")
}
