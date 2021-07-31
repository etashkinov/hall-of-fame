package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/etashkinov/hall-of-fame/controllers"
)

func ApplicationRouter(router *gin.Engine) {
	persons := router.Group("/persons")
	{
		persons.POST("", controllers.CreatePerson)
		persons.GET("", controllers.GetPersons)
		persons.PUT("/:id", controllers.UpdatePerson)
		persons.GET("/:id", controllers.GetPerson)
	}

	dicts := router.Group("/dictionaries")
	{
		dicts.POST("/:type", controllers.CreateDictionary)
		dicts.GET("/:type", controllers.GetDictionaries)
		dicts.PUT("/:type/:id", controllers.UpdateDictionary)
		dicts.DELETE("/:type/:id", controllers.DeleteDictionary)
	}
}
