package services

import (
	"fmt"

	"github.com/Sanket292001/urlshortener/db"
	"github.com/Sanket292001/urlshortener/helper"
	"github.com/Sanket292001/urlshortener/utilities"
	"go.mongodb.org/mongo-driver/bson"
)



func CreateShortenUrlService(urlObj db.UrlBody) (string, error) {	
	// Validate the URL
	if ok, err := utilities.ValidUrl(urlObj.Url); !ok {
		return "", err
	}

	// Check if url is already there
	urlHash := utilities.GenerateHashedString(urlObj.Url)
	hashUrlObj, hashUrlObjErr := db.GetUrlData(bson.M{"origional_hashed_url": urlHash})

	if hashUrlObjErr == nil {
		fmt.Println("Found existing url")
		return hashUrlObj.ShortenUrl, nil
	}
	
	// Generate New Token
	token, err := helper.GenerateToken()
	if err != nil {		
		return "", err
	}

	// Store that in the database
	newUrl := db.Url{
		OrigionalUrl: urlObj.Url,
		OrigionalHashedUrl: urlHash,
		ShortenUrl: token,
		Clicks: 0,		
	}
	_, err = db.InsertNewUrlData(newUrl)
	if err != nil {		
		return "", err
	}

	return token, nil
}