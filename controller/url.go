package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Sanket292001/urlshortener/db"
	"github.com/Sanket292001/urlshortener/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateShortenUrl(w http.ResponseWriter, r *http.Request) {
	var urlObj db.UrlBody

	// Extract the body data
	err := json.NewDecoder(r.Body).Decode(&urlObj)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Generate new url
	token, err := services.CreateShortenUrlService(urlObj)
	
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}


	// Convert the schema to JSON
	url := db.UrlBody{
		Url: "http://localhost:8080/"+token,
	}
	jsonUrlObj, err := json.Marshal(url)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonUrlObj)
}

func GetShortenUrl(w http.ResponseWriter, r *http.Request) {
	path_vars := mux.Vars(r)
	shorten_url := path_vars["shorten_url"]

	// Validate the input
	if len(shorten_url) != 8 {
		w.Write([]byte("Invalid Shorten Url"))
		return
	}

	// Fetch the data from database
	urlObj, urlObjErr := db.GetUrlData(bson.M{"shorten_url": shorten_url})

	if urlObjErr != nil {
		w.Write([]byte(urlObjErr.Error()))
		return
	}

	// Convert the schema to JSON
	jsonUrlObj, err := json.Marshal(urlObj)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonUrlObj)
}

func ShortenRedirectUrl(w http.ResponseWriter, r *http.Request) {
	path_vars := mux.Vars(r)
	shorten_url := path_vars["shorten_url"]

	// Find the Url data from databse
	urlObj, err := db.GetUrlData(bson.M{"shorten_url": shorten_url})

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, urlObj.OrigionalUrl, http.StatusPermanentRedirect)
}