package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/controllers"
)

func ApplicationRouter(router *gin.Engine) {
	router.GET("", func(c *gin.Context) {})

	api := router.Group("/api")

	persons := api.Group("/persons")
	{
		persons.POST("", controllers.CreatePerson)
		persons.GET("", controllers.GetPersons)
	}

	person := persons.Group("/:id")
	{
		person.PUT("", controllers.UpdatePerson)
		person.GET("", controllers.GetPerson)
		person.DELETE("", controllers.DeletePerson)
	}

	personDicts := person.Group("/:type")
	{
		personDicts.GET("", controllers.GetPersonDicts)
		personDicts.POST("", controllers.AddPersonDict)
	}

	personDict := personDicts.Group("/:dictId")
	{
		personDict.PUT("", controllers.UpdatePersonDict)
		personDict.DELETE("", controllers.DeletePersonDict)
	}

	dicts := api.Group("/dictionaries/:type")
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
