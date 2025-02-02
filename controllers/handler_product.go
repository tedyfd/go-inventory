package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go-inventory/internal/database"
	"go-inventory/models"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (uc *UserController) HandlerCreateProduct(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name       string    `json:"name"`
		Quantity   int       `json:"quantity"`
		CategoryID uuid.UUID `json:"category_id"`
		SellerID   uuid.UUID `json:"seller_id"`
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
		SellerID:   params.SellerID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create Product: ", err))
		return
	}

	respondWithJSON(w, 201, "Create product success", models.DatabaseProductToProduct(product))
}

func (uc *UserController) HandlerGetProduct(w http.ResponseWriter, r *http.Request, user database.User) {
	product, _ := uc.Config.DB.GetProduct(r.Context())
	respondWithJSON(w, 200, "Success", models.DatabaseProductsToProducts(product))
}

func (uc *UserController) HandlerDeleteProduct(w http.ResponseWriter, r *http.Request, user database.User) {
	ProductIDStr := chi.URLParam(r, "productID")

	productID, err := uuid.Parse(ProductIDStr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse product ID: ", err))
		return
	}

	err = uc.Config.DB.DeleteProduct(r.Context(), productID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't delete product: ", err))
		return
	}
	respondWithJSON(w, 204, "Delete product success", struct{}{})
}
