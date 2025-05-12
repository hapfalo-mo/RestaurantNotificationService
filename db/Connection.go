package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

func ConnectionDB() (*mongo.Client, error) {
	clientsOption := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientsOption)
	if err != nil {
		log.Fatalf("Error happening in DB Connection progress.. %d", err.Error())
		return nil, err
	}
	// Ping Error Db
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Error happening in DB Connection progress.. %d", err.Error())
		return nil, err
	}
	MongoClient = client
	return client, nil
}

// Get Collection Notification
func GetCollection(name string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatalf("Error happening in DB Connection progress.")
	}
	return MongoClient.Database("restaurantnotification").Collection("notifications")
}
