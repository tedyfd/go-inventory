package controllers

import "net/http"

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "Success", struct{}{})
}
