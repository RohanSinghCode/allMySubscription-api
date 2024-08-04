package main

import (
	"fmt"

	"github.com/RohanSinghCode/allMySubscription-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	//Register route modules
	routes.UserRoutes(r)

	err := godotenv.Load()
	if err != nil {
		fmt.Print("Failed to load environement file")
	}

	r.Run()
}
