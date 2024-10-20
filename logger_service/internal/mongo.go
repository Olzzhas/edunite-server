package internal

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoClient struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoClient(uri string) *MongoClient {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	return &MongoClient{
		client: client,
		db:     client.Database("logs"),
	}
}

func (m *MongoClient) SaveLog(logEntry map[string]interface{}) {
	collection := m.db.Collection("logs")
	_, err := collection.InsertOne(context.Background(), logEntry)
	if err != nil {
		log.Printf("Failed to save log: %v", err)
	}
	log.Println("Log saved successfully")
}
