package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}

	type errResponse struct {
		Error      string `json:"error"`
		StatusCode int    `json:"status_code"`
	}

	respondWithJSON(w, code, errResponse{
		Error:      msg,
		StatusCode: code,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
