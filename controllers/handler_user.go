package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go-inventory/internal/database"
	"go-inventory/models"

	"github.com/google/uuid"
)

func (uc *UserController) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := uc.Config.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username:  params.Username,
		Password:  params.Password,
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create User: %v", err))
		return
	}

	respondWithJSON(w, 201, "Create user success", models.DatabaseUserToUser(user))
}

func (uc *UserController) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, "Success", models.DatabaseUserToUser(user))
}
