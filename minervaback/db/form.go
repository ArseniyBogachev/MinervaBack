package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (f *Form) Save() error {
	if f.ID.IsZero() {
		collection := client.Database("storage").Collection("forms")

		f.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(context.Background(), f)

		return err
	} else {
		return f.Update()
	}
}

func (f *Form) LoadById(id primitive.ObjectID) error {
	collection := client.Database("storage").Collection("forms")

	filter := bson.D{{"_id", id}}

	err := collection.FindOne(context.Background(), filter).Decode(f)

	return err
}

func (f *Form) Update() error {
	collection := client.Database("storage").Collection("forms")

	filter := bson.D{{"_id", f.ID}}
	update := bson.D{{"$set", f}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (f *Form) Delete() error {
	collection := client.Database("storage").Collection("forms")

	filter := bson.D{{"_id", f.ID}}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	collection = client.Database("storage").Collection("formBlocks-" + f.ID.Hex())

	_, err = collection.DeleteMany(context.Background(), bson.D{})

	return err
}

func LoadForms(ownerId primitive.ObjectID) ([]Form, error) {
	collection := client.Database("storage").Collection("forms")

	filter := bson.D{{"owner_id", ownerId}}

	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var forms []Form

	err = cursor.All(context.Background(), &forms)

	if err != nil {
		return nil, err
	}

	return forms, nil
}

func LoadAllForms() ([]Form, error) {
	collection := client.Database("storage").Collection("forms")

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var forms []Form

	err = cursor.All(context.Background(), &forms)

	if err != nil {
		return nil, err
	}

	return forms, nil
}
