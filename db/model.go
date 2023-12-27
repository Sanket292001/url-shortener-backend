package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Url struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OrigionalUrl       string `json:"origionalUrl" bson:"origional_url,omitempty"`
	OrigionalHashedUrl string `json:"origionalHashedUrl" bson:"origional_hashed_url,omitempty"`
	ShortenUrl         string `json:"shortenUrl" bson:"shorten_url,omitempty"`
	Clicks             int `json:"clicks" bson:"clicks,omitempty"`
}

type UrlBody struct {
	Url string `json:"url,omitempty"`
}
