package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/tedyfd/go-inventory/internal/database"
)

func (apiCfg *apiConfig) handlerCreateCategory(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
		return
	}

	description := sql.NullString{}
	if params.Description != "" {
		description.String = params.Description
		description.Valid = true
	}

	category, err := apiCfg.DB.CreateCategory(r.Context(), database.CreateCategoryParams{
		ID:          uuid.New(),
		Name:        params.Name,
		Description: description,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create Category: ", err))
	}

	respondWithJSON(w, 201, databaseCategoryToCategory(category))
}

func (apiCfg *apiConfig) handlerDeleteCategory(w http.ResponseWriter, r *http.Request, user database.User) {
	CategoryIDStr := chi.URLParam(r, "categoryID")

	categoryID, err := uuid.Parse(CategoryIDStr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse category ID: ", err))
		return
	}

	err = apiCfg.DB.DeleteCategory(r.Context(), categoryID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't delete category: ", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
