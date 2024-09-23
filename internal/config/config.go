package config

import "go-inventory/internal/database"

type ApiConfig struct {
	DB *database.Queries
}
