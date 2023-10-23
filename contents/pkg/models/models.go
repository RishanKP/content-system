package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Content struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Story       string             `json:"story" bson:"story,omitempty"`
	UserID      string             `json:"userID" bson:"userID"`
	PublishedOn time.Time          `json:"publishedOn" bson:"publishedOn"`
	UpdatedOn   time.Time          `json:"updatedOn" bson:"updatedOn"`
}

type InsertContentInput struct {
	Title  string `json:"title"`
	Story  string `json:"story"`
	UserID string `json:"userID"`
}

type User struct {
	User struct {
		ID primitive.ObjectID `json:"ID"`
	} `json:"user"`
}

//for parsing interactions api response
type TopContents struct {
	Contents []struct {
		StoryID string `json:"storyID"`
		Likes   int    `json:"likes"`
		Reads   int    `json:"reads"`
	} `json:"contents"`
}

type GetTopContents struct {
	Content Content `json:"content"`
	Likes   int     `json:"likes"`
	Reads   int     `json:"reads"`
}

// --- solely used for swagger documentation purpose
type GetAllContentsResponse struct {
	Message  string    `json:"message" example:"contents fetched"`
	Status   string    `json:"status" example:"success"`
	Contents []Content `json:"contents"`
}

type InsertContentResponse struct {
	Status  string `json:"status" example:"success"`
	Id      string `json:"id" example:"qrewje8234234rwmdsf"`
	Message string `json:"message" example:"content action failed"`
}

type GetContentResponse struct {
	Status  string  `json:"status" example:"success"`
	Message string  `json:"message" example:"content fetched"`
	Content Content `json:"content"`
}

type InternalServerErrorResponse struct {
	Status string `json:"status" example:"failed"`
	Error  string `json:"error" example:"content action failed"`
}

type GetTopContentsResponse struct {
	Message  string           `json:"message" example:"contents fetched"`
	Status   string           `json:"status" example:"success"`
	Contents []GetTopContents `json:"contents"`
}

//only used for swagger documentation purpose ----
