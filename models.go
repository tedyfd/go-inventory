package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/tedyfd/go-inventory/internal/database"
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

func databaseUserToUser(dbUser database.User) User {
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

func databaseCategoryToCategory(dbCategory database.Category) Category {
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

func databaseProductToProduct(dbProduct database.Product) Product {
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
