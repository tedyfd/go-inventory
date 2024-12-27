package controllers

import (
	"encoding/json"
	"fmt"
	"go-inventory/internal/database"
	"go-inventory/models"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (uc *UserController) HandlerCreateOrder(w http.ResponseWriter, r *http.Request, user database.User) {
	type OrderDetail struct {
		ProductID string `json:"product_id"`
		Quantity  int32  `json:"quantity"`
	}

	type Order struct {
		CustomerID  string        `json:"customer_id"`
		OrderDetail []OrderDetail `json:"order_detail"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Order{}
	errD := decoder.Decode(&params)
	if errD != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing r.Body JSON: ", errD))
		fmt.Printf("Error parsing r.Body JSON:")
		return
	}

	customerID, err := uuid.Parse(params.CustomerID)
	order, err := uc.Config.DB.CreateOrder(r.Context(), database.CreateOrderParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		CustomerID: customerID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error add order: ", err))
		return
	}

	for _, detail := range params.OrderDetail {
		productID, _ := uuid.Parse(detail.ProductID)
		_, err := uc.Config.DB.CreateOrderDetailAndUpdateProduct(r.Context(), database.CreateOrderDetailAndUpdateProductParams{
			POrderID:   order.ID,
			PProductID: productID,
			PQuantity:  detail.Quantity,
		})
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error add order detail: ", err))
			return
		}
	}
	respondWithJSON(w, 200, "success create order", models.DatabaseOrderToOrder(order))
}
