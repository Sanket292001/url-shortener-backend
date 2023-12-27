package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionUrl = ""
const databaseName = "urlshortener"

var collections = make(map[string]*mongo.Collection)

func InitDatabase() {
	opts := options.Client().ApplyURI(connectionUrl)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}
	
	collections["url"] = client.Database(databaseName).Collection("url")
}

func GetCollectionReference(collectionName string) *mongo.Collection {
	return collections[collectionName]
}