package controllers

import (
	"net/http"

	"github.com/RohanSinghCode/allMySubscription-api/models/request"
	"github.com/RohanSinghCode/allMySubscription-api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUser(c *gin.Context) {
	var uri struct {
		UserId string `uri:"id"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := uuid.Parse(uri.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
	}
	user, err := services.GetUser(userId, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user"})
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
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
