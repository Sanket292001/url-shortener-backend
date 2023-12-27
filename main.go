package main

import (
	"log"
	"net/http"

	"github.com/Sanket292001/urlshortener/db"
	"github.com/Sanket292001/urlshortener/routes"
)

func main() {
	router := routes.AddRoutesAndMiddlewares()
	
	db.InitDatabase()

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}