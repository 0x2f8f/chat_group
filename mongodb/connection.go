package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"log"
)

const db_url = "mongodb://127.0.0.1:27017"
const db_name = "chats"

func getClientDb() *mongo.Client {
	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(db_url))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetCollection(collection_name string) *mongo.Collection {

	return getClientDb().Database(db_name).Collection(collection_name)
}