package routes

import (
	"net/http"

	"github.com/Sanket292001/urlshortener/controller"
	"github.com/Sanket292001/urlshortener/utilities"
	"github.com/gorilla/mux"
)

func addUrlRoutes(router *mux.Router) {
	router.HandleFunc(utilities.POST_CREATE_SHORTEN_ENDPOINT, controller.CreateShortenUrl).Methods(http.MethodPost)
	router.HandleFunc(utilities.GET_SHORTEN_ENDPOINT, controller.GetShortenUrl).Methods(http.MethodGet)
	router.HandleFunc(utilities.GET_SHORTEN_REDIRECT_ENDPOINT, controller.ShortenRedirectUrl).Methods(http.MethodGet)
}