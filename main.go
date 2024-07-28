package main

import (
	"github.com/RohanSinghCode/allMySubscription-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Register route modules
	routes.UserRoutes(r)

	r.Run()
}
