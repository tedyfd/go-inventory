package controllers

import "go-inventory/internal/config"

type UserController struct {
	Config *config.ApiConfig // Use 'Config' or 'ApiConfig' depending on your naming convention
}
