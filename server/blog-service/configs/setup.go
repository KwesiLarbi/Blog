package configs

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// creating new client based on uri
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	// context with timeout is mainly used when we want to make an
	// external request, such as a database request
	// create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	return client
}

// client database instance
var DB *mongo.Client = ConnectDB()

// GetCollection is a function makes a connection with a collection in the database
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("blog").Collection(collectionName)
	
	return collection
}
