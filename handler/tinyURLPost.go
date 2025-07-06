package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type URLBody struct {
	URL string `json:"url"`
}

func (h *Handler) PostURLToGetTinyURL(w http.ResponseWriter, r *http.Request) {
	url := URLBody{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&url)
	if err != nil {
		log.Println("error occurred while decoding request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tinyurl := h.Services.TinyURL.ConvertURLToTinyURL(url.URL)

	w.Header().Set("Content-Type", "application/json")

	log.Printf("%+v")
	err = json.NewEncoder(w).Encode(tinyurl)
	if err != nil {
		fmt.Println("Error encoding response:", err)
		return
	}
}
