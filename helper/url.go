package helper

import (
	"fmt"

	"github.com/Sanket292001/urlshortener/db"
	"github.com/Sanket292001/urlshortener/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateToken() (string, error) {
	noDocuments := false
	var token string

	// Loop until new token found in the database
	for !noDocuments {
		fmt.Println("Generating Shorten Url")

		// Generate the Random ID
		token = utilities.RandomToken(8)

		// Check if its a unique id
		_, err := db.GetUrlData(bson.M{"shorten_url": token})

		if err != nil {
			if err == mongo.ErrNoDocuments {
				fmt.Println("No Documents Found")
				noDocuments = true
			} else {
				return "", err
			}
		}
	}

	return token, nil
}