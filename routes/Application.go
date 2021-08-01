package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/controllers"
)

func ApplicationRouter(router *gin.Engine) {
	router.GET("", func(c *gin.Context) {})

	persons := router.Group("/persons")
	{
		persons.POST("", controllers.CreatePerson)
		persons.GET("", controllers.GetPersons)
	}

	person := persons.Group("/:id")
	{
		person.PUT("", controllers.UpdatePerson)
		person.GET("", controllers.GetPerson)
	}

	skills := person.Group("/skills")
	{
		skills.GET("", controllers.GetPersonSkills)
		skills.POST("", controllers.AddPersonSkill)
	}

	dicts := router.Group("/dictionaries/:type")
	{
		dicts.POST("", controllers.CreateDictionary)
		dicts.GET("", controllers.GetDictionaries)
	}

	dict := dicts.Group("/:id")
	{
		dict.PUT("", controllers.UpdateDictionary)
		dict.DELETE("", controllers.DeleteDictionary)
	}
}
