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

func (uc *UserController) HandlerCreateCustomer(w http.ResponseWriter, r *http.Request, user database.User) {
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

	customer, err := uc.Config.DB.CreateCustomer(r.Context(), database.CreateCustomerParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create customer: ", err))
		return
	}

	respondWithJSON(w, 201, "Create customer success", models.DatabaseCustomerToCustomer(customer))
}

func (uc *UserController) HandlerGetCustomer(w http.ResponseWriter, r *http.Request, user database.User) {
	customer, _ := uc.Config.DB.GetCustomer(r.Context())
	respondWithJSON(w, 200, "Success", models.DatabaseCustomersToCustomers(customer))
}

func (uc *UserController) HandlerGetCustomerByName(w http.ResponseWriter, r *http.Request, user database.User) {
	customerName := chi.URLParam(r, "customerName")

	customer, err := uc.Config.DB.GetCustomerByName(r.Context(), customerName)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error get customer: ", err))
		return
	}
	respondWithJSON(w, 200, "Get customer success", models.DatabaseCustomerToCustomer(customer))
}

func (uc *UserController) HandlerUpdateCustomer(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}
	customerID, err := uuid.Parse(chi.URLParam(r, "customerID"))
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	errD := decoder.Decode(&params)
	if errD != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing r.Body JSON:"))
		fmt.Printf("Error parsing r.Body JSON:")
		return
	}

	customer, err := uc.Config.DB.UpdateCustomer(r.Context(), database.UpdateCustomerParams{
		Name:      params.Name,
		UpdatedAt: time.Now(),
		ID:        customerID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error update customer: ", err))
		return
	}
	respondWithJSON(w, 200, "Success update customer", models.DatabaseCustomerToCustomer(customer))
}

func (uc *UserController) HandlerDeleteCustomer(w http.ResponseWriter, r *http.Request, user database.User) {
	CustomerIDStr := chi.URLParam(r, "customerID")

	customerID, err := uuid.Parse(CustomerIDStr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse customer ID: ", err))
		return
	}

	err = uc.Config.DB.DeleteCustomer(r.Context(), customerID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't delete category: ", err))
		return
	}
	respondWithJSON(w, 204, "Delete Seller success", struct{}{})
}
