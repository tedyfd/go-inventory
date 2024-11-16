package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int       `json:"status_code"`
	Error      []Message `json:"error"`
}

type Message struct {
	Message string `json:"message"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}

	response := ErrorResponse{
		StatusCode: code,
		Error: []Message{
			{Message: msg},
		},
	}

	sendResponse(w, code, response)
}

func respondWithJSON(w http.ResponseWriter, code int, msg string, payload interface{}) {
	response := JsonResponse{
		StatusCode: code,
		Message:    msg,
		Data:       payload,
	}

	sendResponse(w, code, response)
}

func sendResponse(w http.ResponseWriter, code int, response interface{}) {
	dat, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed to marshal JSON response %v", response)
		w.WriteHeader(500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
