package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var client *mongo.Client

func init() {
	Connect()
}

func Connect() {
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/", mongoUser, mongoPassword, mongoHost, mongoPort)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		panic(err)
	}
}
