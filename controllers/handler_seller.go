package controllers

import (
	"encoding/json"
	"fmt"
	"go-inventory/internal/database"
	"go-inventory/models"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (uc *UserController) HandlerCreateSeller(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
		return
	}

	seller, err := uc.Config.DB.CreateSeller(r.Context(), database.CreateSellerParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create seller: ", err))
		return
	}

	respondWithJSON(w, 201, "success create seller", seller)
}

func (uc *UserController) HandlerGetSeller(w http.ResponseWriter, r *http.Request, user database.User) {
	sellers, _ := uc.Config.DB.GetSeller(r.Context())
	respondWithJSON(w, 200, "Success get all seller", models.DatabaseSellersToSellers(sellers))
}

func (uc *UserController) HandlerGetSellerByName(w http.ResponseWriter, r *http.Request, user database.User) {
	sellerName := chi.URLParam(r, "sellerName")
	seller, err := uc.Config.DB.GetSellerByName(r.Context(), sellerName)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get seller: ", err))
		return
	}
	respondWithJSON(w, 200, "Success get seller by name", models.DatabaseSellerToSeller(seller))
}

func (uc *UserController) HandlerUpdateSeller(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}
	sellerID, err := uuid.Parse(chi.URLParam(r, "sellerID"))
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

	seller, err := uc.Config.DB.UpdateSeller(r.Context(), database.UpdateSellerParams{
		Name:      params.Name,
		UpdatedAt: time.Now(),
		ID:        sellerID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Update Seller: ", err))
		return
	}
	respondWithJSON(w, 200, "Success update seller", models.DatabaseSellerToSeller(seller))
}
