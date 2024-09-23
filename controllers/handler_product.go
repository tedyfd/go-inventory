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

func (uc *UserController) HandlerCreateProduct(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name       string    `json:"name"`
		Quantity   int       `json:"quantity"`
		CategoryID uuid.UUID `json:"category_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: ", err))
		return
	}

	product, err := uc.Config.DB.CreateProduct(r.Context(), database.CreateProductParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Name:       params.Name,
		Quantity:   int32(params.Quantity),
		UserID:     user.ID,
		CategoryID: params.CategoryID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create Product: ", err))
	}

	respondWithJSON(w, 201, models.DatabaseProductToProduct(product))
}
