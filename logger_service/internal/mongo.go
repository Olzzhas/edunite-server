package internal

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectMongoDB(uri string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client.Database("logs"), nil
}

func SaveLogToMongo(db *mongo.Database, logData interface{}) error {
	collection := db.Collection("logs")
	_, err := collection.InsertOne(context.Background(), logData)
	if err != nil {
		log.Printf("Failed to save log to MongoDB: %v", err)
		return err
	}
	return nil
}
