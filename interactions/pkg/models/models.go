package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Interactions struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	StoryID string             `json:"storyID" bson:"storyID"`
	Likes   int                `json:"likes" bson:"likes"`
	Reads   int                `json:"reads" bson:"reads"`
	LikedBy []string           `json:"likedBy" bson:"likedBy"`
	ReadBy  []string           `json:"readBy" bson:"readBy"`
}

type InteractionsInput struct {
	StoryID string `json:"storyID"`
	UserID  string `json:"userID"`
	Type    string `json:"type"`
}

type User struct {
	User struct {
		ID primitive.ObjectID `json:"ID"`
	} `json:"user"`
}

//------------ solely for swagger documentation purpose
type GetInteractionsByStoryID struct {
	Message      string       `json:"message" example:"interactions fetched"`
	Status       string       `json:"status" example:"success"`
	Interactions Interactions `json:"interactions"`
}

type InteractionResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:" action completed"`
}

type InternalServerErrorResponse struct {
	Status string `json:"status" example:"failed"`
	Error  string `json:"error" example:"action failed"`
}

type GetTopInteractionsResponse struct {
	Status   string         `json:"status" example:"success"`
	Message  string         `json:"message" example:" action completed"`
	Contents []Interactions `json:"contents"`
}

// solely for swagger documentation purpose -----------
