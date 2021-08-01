package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/persistence"
)

type skillRequest struct {
	SkillId int64
	Since   string
	Level   persistence.Level
}

func AddPersonSkill(c *gin.Context) {
	personId, _ := getId(c)

	var request skillRequest
	bindJSON(c, &request)

	var since persistence.Date
	since.Parse(request.Since)
	skill, err := persistence.AddPersonSkill(personId, request.SkillId, since, request.Level)

	ok(c, skill, err)
}

func GetPersonSkills(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		serverError(c, err)
		return
	}

	skills, err := persistence.GetPersonSkills(id)

	ok(c, skills, err)
}
