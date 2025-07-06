package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (h *Handler) RedirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	// log.Println("code was here")
	// vars := mux.Vars(r)
	// path := vars["any"]

	// originalURL := h.Datalayer.TinyURLDatabase.GetOriginalURL(path)

	// log.Println("code was here", path, originalURL)
	// w.Write([]byte("originalURL: " + originalURL))

	response := MyResponse{
		Message: "Hello from the server!",
		Code:    200,
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("Error encoding response:", err)
		return
	}
}
