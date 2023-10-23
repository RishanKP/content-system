package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	LastName string             `json:"lastName" bson:"lastname,omitempty"`
	Email    string             `json:"email" bson:"email"`
	Phone    string             `json:"phone" bson:"phone"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type GetAllUsersResponse struct {
	Message string `json:"message" example:"users fetched"`
	Status  string `json:"status" example:"success"`
	Users   []User `json:"users"`
}

type CreateUserResponse struct {
	Status  string `json:"status" example:"success"`
	Id      string `json:"id" example:"qrewje8234234rwmdsf"`
	Message string `json:"message" example:"user action completed"`
}

type GetUserResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"user data fetched"`
	User    User   `json:"user"`
}

type InternalServerErrorResponse struct {
	Status string `json:"status" example:"failed"`
	Error  string `json:"error" example:"user action failed"`
}
