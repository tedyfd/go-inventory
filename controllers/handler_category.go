package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-inventory/internal/database"
	"go-inventory/models"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (uc *UserController) HandlerCreateCategory(w http.ResponseWriter, r *http.Request, user database.User) {
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

	category, err := uc.Config.DB.CreateCategory(r.Context(), database.CreateCategoryParams{
		ID:          uuid.New(),
		Name:        params.Name,
		Description: description,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create Category: ", err))
		return
	}

	respondWithJSON(w, 201, "Create category success", models.DatabaseCategoryToCategory(category))
}

func (uc *UserController) HandlerGetCategory(w http.ResponseWriter, r *http.Request, user database.User) {
	categories, _ := uc.Config.DB.GetCategory(r.Context())
	respondWithJSON(w, 200, "Success get all categories", models.DatabasecategoriesTocategories(categories))
}

func (uc *UserController) HandlerGetCategoryByName(w http.ResponseWriter, r *http.Request, user database.User) {
	categoryName := chi.URLParam(r, "categoryName")

	category, err := uc.Config.DB.GetCategoryByName(r.Context(), categoryName)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error get category: ", err))
		return
	}
	respondWithJSON(w, 200, "Get category success", models.DatabaseCategoryToCategory(category))
}

func (uc *UserController) HandlerUpdateCategory(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	categoryID, err := uuid.Parse(chi.URLParam(r, "categoryID"))
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	errD := decoder.Decode(&params)
	if errD != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing r.Body JSON:", err))
		fmt.Printf("Error parsing r.Body JSON:", err)
		return
	}

	description := sql.NullString{}
	if params.Description != "" {
		description.String = params.Description
		description.Valid = true
	}

	category, err := uc.Config.DB.UpdateCategory(r.Context(), database.UpdateCategoryParams{
		Name:        params.Name,
		Description: description,
		ID:          categoryID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Update Category: ", err))
		return
	}
	respondWithJSON(w, 200, "Success update Category", models.DatabaseCategoryToCategory(category))
}

func (uc *UserController) HandlerDeleteCategory(w http.ResponseWriter, r *http.Request, user database.User) {
	CategoryIDStr := chi.URLParam(r, "categoryID")

	categoryID, err := uuid.Parse(CategoryIDStr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse category ID: ", err))
		return
	}

	err = uc.Config.DB.DeleteCategory(r.Context(), categoryID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't delete category: ", err))
		return
	}
	respondWithJSON(w, 204, "Delete category success", struct{}{})
}
