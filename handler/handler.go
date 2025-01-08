package handler

import (
	"log"
	"net/http"
	"url-shortner/datalayer"
	"url-shortner/services"

	"github.com/gorilla/mux"
)

type HandlerInterface interface {
	PostURLToGetTinyURL(w http.ResponseWriter, r *http.Request)
	RedirectToOriginalURL(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Services  services.Services
	Datalayer datalayer.Datalayer
}

func NewHandlerInterface(urlMap map[string]string) HandlerInterface {
	return &Handler{
		Services:  services.InitServices(urlMap),
		Datalayer: datalayer.InitDatalayer(urlMap),
	}
}

func StartController(urlMap map[string]string) {
	log.Println("code was here")
	h := NewHandlerInterface(urlMap)

	r := mux.NewRouter()
	r.HandleFunc("/generatetinyurl", h.PostURLToGetTinyURL)
	// r.HandleFunc("/redirect/{any}", h.RedirectToOriginalURL)

	log.Println("running on localhost:8080")
	http.ListenAndServe(":8080", r)
}
