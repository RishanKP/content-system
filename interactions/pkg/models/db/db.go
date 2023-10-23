package db

import (
	"content-system/interactions/pkg/models"
	"context"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InteractionsModel struct {
	Client *mongo.Collection
}

func (m *InteractionsModel) GetInteractionsByStoryID(id string) (models.Interactions, error) {

	var u models.Interactions
	err := m.Client.FindOne(context.TODO(), bson.M{"storyID": id}).Decode(&u)

	if err != nil {
		return models.Interactions{}, err
	}

	return u, nil
}

func (m *InteractionsModel) Save(interaction models.InteractionsInput, actionType string) (*mongo.UpdateResult, error) {

	opts := options.Update().SetUpsert(true)

	var field1, field2 string

	if interaction.Type == "read" {
		field1 = "reads"
		field2 = "readBy"
	} else if interaction.Type == "like" {
		field1 = "likes"
		field2 = "likedBy"
	} else {
		return nil, errors.New("invalid type. supported types [read,like]")
	}

	action := "$addToSet"
	incrementValue := 1
	condition := "$nin"

	if actionType == "remove" {
		incrementValue = -1
		action = "$pull"
		condition = "$in"
	}

	filter := bson.M{
		"storyID": interaction.StoryID,
		field2: bson.M{
			condition: []interface{}{interaction.UserID},
		},
	}

	update := bson.M{
		"$inc": bson.M{
			field1: incrementValue,
		},
		action: bson.M{
			field2: interaction.UserID,
		},
	}

	return m.Client.UpdateOne(context.TODO(), filter, update, opts)
}

func (m *InteractionsModel) GetTopContents(param, limit string) ([]models.Interactions, error) {
	contents := []models.Interactions{}

	lmt, err := strconv.Atoi(limit)
	if err != nil {
		lmt = 10
	}

	if param != "reads" && param != "likes" {
		return contents, errors.New("invalid parameter. supported parameters are reads and likes")
	}

	opts := options.Find()
	opts.SetSort(bson.M{param: -1})
	opts.SetLimit(int64(lmt))

	cursor, err := m.Client.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return contents, err
	}

	err = cursor.All(context.TODO(), &contents)
	if err != nil {
		return contents, err
	}

	return contents, nil
}

func getObjectID(id string) (primitive.ObjectID, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return _id, nil
}
