package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUrlData(filter bson.M) (Url, error) {
	var urlCollection = GetCollectionReference("url")
	
	var urlObj Url
	err := urlCollection.FindOne(context.Background(), filter).Decode(&urlObj)
	if err != nil {
		return urlObj, err
	}

	return urlObj, nil
}

func InsertNewUrlData(urlObj Url) (interface{}, error) {
	var urlCollection = GetCollectionReference("url")
	
	inserted, err := urlCollection.InsertOne(context.Background(), urlObj)
	return inserted.InsertedID, err
}