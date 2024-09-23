package controllers

import (
	"fmt"
	"net/http"

	"go-inventory/internal/auth"
	"go-inventory/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (uc *UserController) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %w", err))
			return
		}

		user, err := uc.Config.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %w", err))
			return
		}

		handler(w, r, user)
	}
}
