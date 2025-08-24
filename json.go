package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Printf("Server error: %s", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}

	jsonResponse(w, code, errorResponse{Error: message})
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
