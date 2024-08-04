package controllers

import (
	"net/http"

	"github.com/RohanSinghCode/allMySubscription-api/models/request"
	"github.com/RohanSinghCode/allMySubscription-api/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Setup Done"})
}

func InsertUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := services.InsertUser(userRequest, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": userId})
}
