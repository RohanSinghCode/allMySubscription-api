package routes

import (
	"github.com/RohanSinghCode/allMySubscription-api/controllers"
	"github.com/gin-gonic/gin"
)

func ProductsRoute(router *gin.Engine) {
	productGroup := router.Group("users/:userId/products")
	{
		productGroup.POST("/", controllers.InsertUser)
		productGroup.GET("/:productId", controllers.GetUser)
		productGroup.GET("/")
	}
}
