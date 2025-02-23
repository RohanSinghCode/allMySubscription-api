package services

import (
	"database/sql"
	"errors"

	"github.com/RohanSinghCode/allMySubscription-api/internal/database"
	"github.com/RohanSinghCode/allMySubscription-api/models/request"
	"github.com/RohanSinghCode/allMySubscription-api/models/response"
	"github.com/RohanSinghCode/allMySubscription-api/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InsertUser(userRequest request.UserRequest, c *gin.Context) (uuid.UUID, error) {
	invalidError := validateRequest(userRequest)

	if invalidError != nil {
		return uuid.Nil, invalidError
	}

	postgres := repository.Connection{}
	conn, err := postgres.ConnectToDb()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()
	user, err := postgres.DB.InsertUser(c, database.InsertUserParams{
		Firstname: userRequest.FirstName,
		Lastname:  sql.NullString{String: userRequest.LastName, Valid: true},
		Email:     userRequest.Email,
		Password:  userRequest.Password,
	})

	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func GetUser(userId uuid.UUID, c *gin.Context) (*response.UserResponse, error) {
	if userId == uuid.Nil {
		return nil, errors.New("invalid user id")
	}

	postgres := repository.Connection{}
	conn, err := postgres.ConnectToDb()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	user, err := postgres.DB.GetUser(c, userId)

	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		Id:        user.ID,
		FirstName: user.Firstname,
		LastName:  user.Lastname.String,
		Email:     user.Email,
	}, nil
}

func validateRequest(userRequest request.UserRequest) error {
	if userRequest.FirstName == "" {
		return errors.New("first Name is required")
	}

	if userRequest.Email == "" {
		return errors.New("email is required")
	}

	if userRequest.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
