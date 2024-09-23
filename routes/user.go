package routes

import (
	"github.com/RohanSinghCode/allMySubscription-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", controllers.InsertUser)
		userGroup.GET("/:id", controllers.GetUser)
	}
}
