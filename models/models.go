package models

import (
	"time"

	"go-inventory/internal/database"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Category struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func DatabaseCategoryToCategory(dbCategory database.Category) Category {
	var description *string
	if dbCategory.Description.Valid {
		description = &dbCategory.Description.String
	}
	return Category{
		ID:          dbCategory.ID,
		Name:        dbCategory.Name,
		Description: description,
	}
}

type Product struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	Quantity   int       `json:"quantity"`
	UserID     uuid.UUID `json:"user_id"`
	categoryID uuid.UUID `json:"category_id"`
}

func DatabaseProductToProduct(dbProduct database.Product) Product {
	return Product{
		ID:         dbProduct.ID,
		CreatedAt:  dbProduct.CreatedAt,
		UpdatedAt:  dbProduct.UpdatedAt,
		Name:       dbProduct.Name,
		Quantity:   int(dbProduct.Quantity),
		UserID:     dbProduct.UserID,
		categoryID: dbProduct.CategoryID,
	}
}

type Customer struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func DatabaseCustomerToCustomer(dbCustomer database.Customer) Customer {
	return Customer{
		ID:        dbCustomer.ID,
		CreatedAt: dbCustomer.CreatedAt,
		UpdatedAt: dbCustomer.UpdatedAt,
		Name:      dbCustomer.Name,
	}
}

func DatabaseCustomersToCustomers(dbCustomer []database.Customer) []Customer {
	transform := make([]Customer, len(dbCustomer))
	for i, customer := range dbCustomer {
		transform[i] = Customer{
			ID:        customer.ID,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
			Name:      customer.Name,
		}
	}
	return transform
}

type Seller struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func DatabaseSellersToSellers(dbSeller []database.Seller) []Seller {
	transform := make([]Seller, len(dbSeller))
	for i, seller := range dbSeller {
		transform[i] = Seller{
			ID:        seller.ID,
			CreatedAt: seller.CreatedAt,
			UpdatedAt: seller.UpdatedAt,
			Name:      seller.Name,
		}
	}
	return transform
}

func DatabaseSellerToSeller(dbSeller database.Seller) Seller {
	return Seller{
		ID:        dbSeller.ID,
		CreatedAt: dbSeller.CreatedAt,
		UpdatedAt: dbSeller.UpdatedAt,
		Name:      dbSeller.Name,
	}
}

type Order struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseOrderToOrder(dbOrder database.Order) Order {
	return Order{
		ID:        dbOrder.ID,
		CreatedAt: dbOrder.CreatedAt,
		UpdatedAt: dbOrder.UpdatedAt,
	}
}
