package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB create client failed. Error: ", err.Error())
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MongoDB connection failed. Error: ", err.Error())
	}
	return client
}
