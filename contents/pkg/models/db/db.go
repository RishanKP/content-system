package db

import (
	"content-system/contents/pkg/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContentModel struct {
	Client *mongo.Collection
}

func (m *ContentModel) All() ([]models.Content, error) {
	contents := []models.Content{}

	cursor, err := m.Client.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &contents)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (m *ContentModel) Latest() ([]models.Content, error) {
	contents := []models.Content{}

	opts := options.Find()
	opts.SetSort(bson.M{"publishedOn": -1})

	cursor, err := m.Client.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &contents)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (m *ContentModel) GetContentById(id string) (models.Content, error) {
	_id, err := getObjectID(id)
	if err != nil {
		return models.Content{}, err
	}
	var c models.Content
	err = m.Client.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&c)

	if err != nil {
		return models.Content{}, err
	}

	return c, nil
}

func (m *ContentModel) Insert(content models.InsertContentInput) (*mongo.InsertOneResult, error) {
	c := models.Content{
		UserID:      content.UserID,
		Title:       content.Title,
		Story:       content.Story,
		PublishedOn: time.Now(),
	}

	return m.Client.InsertOne(context.TODO(), c)
}

func (m *ContentModel) Update(id string, content models.InsertContentInput) (*mongo.UpdateResult, error) {

	_id, err := getObjectID(id)
	if err != nil {
		return nil, err
	}

	return m.Client.UpdateOne(context.TODO(), bson.M{"_id": _id}, bson.M{
		"$set": bson.M{
			"title":     content.Title,
			"story":     content.Story,
			"updatedOn": time.Now(),
		},
	})
}

func (m *ContentModel) Delete(id string) (*mongo.DeleteResult, error) {
	ctx := context.TODO()

	_id, err := getObjectID(id)
	if err != nil {
		return nil, err
	}

	return m.Client.DeleteOne(ctx, bson.M{"_id": _id})
}

func (m *ContentModel) GetTopContents(topContents models.TopContents) ([]models.GetTopContents, error) {
	var res []models.GetTopContents
	for i := range topContents.Contents {
		var c models.Content

		_id, err := getObjectID(topContents.Contents[i].StoryID)
		if err != nil {
			//do nothing..parse remaining
		}

		err = m.Client.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&c)

		if err != nil {
			//do nothing..fetch remaining
		}

		res = append(res, models.GetTopContents{
			Content: c,
			Likes:   topContents.Contents[i].Likes,
			Reads:   topContents.Contents[i].Reads,
		})
	}

	return res, nil
}

func getObjectID(id string) (primitive.ObjectID, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return _id, nil
}
