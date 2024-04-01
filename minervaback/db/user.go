package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *User) Save() error {
	if u.ID.IsZero() {
		collection := client.Database("storage").Collection("users")

		u.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(context.Background(), u)

		return err
	} else {
		return u.Update()
	}
}

func (u *User) LoadById(id primitive.ObjectID) error {
	collection := client.Database("storage").Collection("users")

	filter := bson.D{{"_id", id}}

	err := collection.FindOne(context.Background(), filter).Decode(u)

	return err
}

func (u *User) LoadByLogin(login string) error {
	collection := client.Database("storage").Collection("users")

	filter := bson.D{{"login", login}}

	err := collection.FindOne(context.Background(), filter).Decode(u)

	return err
}

func (u *User) Update() error {
	collection := client.Database("storage").Collection("users")

	filter := bson.D{{"_id", u.ID}}
	update := bson.D{{"$set", u}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (u *User) Delete() error {
	collection := client.Database("storage").Collection("users")

	filter := bson.D{{"_id", u.ID}}

	_, err := collection.DeleteOne(context.Background(), filter)

	return err
}

func LoadUsers() ([]*User, error) {
	collection := client.Database("storage").Collection("users")

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	users := make([]*User, 0)

	for cursor.Next(context.Background()) {
		user := &User{}
		err := cursor.Decode(user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
