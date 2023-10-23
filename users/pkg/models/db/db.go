package db

import (
	"content-system/users/pkg/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Client *mongo.Collection
}

func (m *UserModel) All() ([]models.User, error) {
	users := []models.User{}

	cursor, err := m.Client.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserModel) GetUserById(id string) (models.User, error) {
	_id, err := getObjectID(id)
	if err != nil {
		return models.User{}, err
	}
	var u models.User
	err = m.Client.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&u)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (m *UserModel) Create(user models.CreateUserInput) (*mongo.InsertOneResult, error) {

	var u models.User
	err := m.Client.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&u)
	if err != nil {
		//do nothing
	}

	if u.Email == user.Email {
		return nil, errors.New("email Id already exist")
	}

	return m.Client.InsertOne(context.TODO(), user)
}

func (m *UserModel) Update(id string, user models.CreateUserInput) (*mongo.UpdateResult, error) {

	_id, err := getObjectID(id)
	if err != nil {
		return nil, err
	}

	return m.Client.UpdateOne(context.TODO(), bson.M{"_id": _id}, bson.M{
		"$set": bson.M{
			"name":     user.Name,
			"lastName": user.LastName,
			"email":    user.Email,
			"phone":    user.Phone,
		},
	})
}

func (m *UserModel) Delete(id string) (*mongo.DeleteResult, error) {
	ctx := context.TODO()

	_id, err := getObjectID(id)
	if err != nil {
		return nil, err
	}

	return m.Client.DeleteOne(ctx, bson.M{"_id": _id})
}

func getObjectID(id string) (primitive.ObjectID, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return _id, nil
}
