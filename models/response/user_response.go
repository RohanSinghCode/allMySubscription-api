package response

import "github.com/google/uuid"

type UserResponse struct {
	Id        uuid.UUID `json:"Id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}
