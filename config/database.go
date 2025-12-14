package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var MongoClient *mongo.Client

// ConnectDatabase establishes a connection to MongoDB
func ConnectDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(AppConfig.MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
		return err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
		return err
	}

	MongoClient = client
	DB = client.Database(AppConfig.DatabaseName)

	log.Println("Connected to MongoDB successfully!")
	return nil
}

// DisconnectDatabase closes the MongoDB connection
func DisconnectDatabase() error {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := MongoClient.Disconnect(ctx)
		if err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
			return err
		}
		log.Println("Disconnected from MongoDB")
	}
	return nil
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
