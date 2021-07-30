package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/etashkinov/hall-of-fame/controllers"
)

func ApplicationV1Router(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		// Documentation Swagger
		{
			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		// Medicines
		v1Medicines := v1.Group("/persons")
		{
			v1Medicines.POST("", controllers.CreatePerson)
			v1Medicines.GET("", controllers.GetPersons)
			v1Medicines.PUT("/:id", controllers.UpdatePerson)
			v1Medicines.GET("/:id", controllers.GetPerson)
		}
	}
}
