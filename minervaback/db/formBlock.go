package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (f *FormBlock) Save() error {
	if f.ID.IsZero() {
		collection := client.Database("storage").Collection("formBlocks-" + f.FormID.Hex())

		var err error

		f.ID = primitive.NewObjectID()
		f.Order, err = getFormBlocksCount(f.FormID)

		if err != nil {
			fmt.Printf("Failed to get max block order: %v\n", err)
			f.Order = 0
		}

		_, err = collection.InsertOne(context.Background(), f)

		return err
	} else {
		return f.Update()
	}
}

func (f *FormBlock) Load(order int) error {
	collection := client.Database("storage").Collection("formBlocks-" + f.FormID.Hex())

	filter := bson.D{{"order", order}}

	return collection.FindOne(context.Background(), filter).Decode(f)
}

func (f *FormBlock) Update() error {
	collection := client.Database("storage").Collection("formBlocks-" + f.FormID.Hex())

	filter := bson.D{{"_id", f.ID}}
	update := bson.D{{"$set", f}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (f *FormBlock) MoveTo(order int) error {
	size, err := getFormBlocksCount(f.FormID)

	if err != nil || size <= 1 || order == f.Order {
		return err
	}

	if size <= order {
		order = size - 1
	} else if order < 0 {
		order = 0
	}

	collection := client.Database("storage").Collection("formBlocks-" + f.FormID.Hex())

	if order > f.Order {
		_, err = collection.UpdateMany(context.Background(), bson.D{
			{"order", bson.D{{"$gt", f.Order}}},
			{"order", bson.D{{"$lte", order}}},
		}, bson.D{{"$inc", bson.D{{"order", -1}}}})
	} else {
		_, err = collection.UpdateMany(context.Background(), bson.D{
			{"order", bson.D{{"$lt", f.Order}}},
			{"order", bson.D{{"$gte", order}}},
		}, bson.D{{"$inc", bson.D{{"order", 1}}}})
	}

	if err != nil {
		return err
	}

	filter := bson.D{{"_id", f.ID}}
	update := bson.D{{"$set", bson.D{{"order", order}}}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (f *FormBlock) Delete() error {
	collection := client.Database("storage").Collection("formBlocks-" + f.FormID.Hex())

	filter := bson.D{{"_id", f.ID}}

	_, err := collection.DeleteOne(context.Background(), filter)

	return err
}

func getFormBlocksCount(formId primitive.ObjectID) (int, error) {
	collection := client.Database("storage").Collection("formBlocks-" + formId.Hex())

	v, err := collection.CountDocuments(context.Background(), bson.D{})

	return int(v), err
}

func LoadFormBlocks(formId primitive.ObjectID) ([]FormBlock, error) {
	collection := client.Database("storage").Collection("formBlocks-" + formId.Hex())

	// sort by order
	cursor, err := collection.Find(context.Background(), bson.D{}, &options.FindOptions{Sort: bson.D{{"order", 1}}})

	if err != nil {
		return nil, err
	}

	var formBlocks []FormBlock
	err = cursor.All(context.Background(), &formBlocks)

	if err != nil {
		return nil, err
	}

	for i := range formBlocks {
		formBlocks[i].FormID = formId
	}

	return formBlocks, nil
}
