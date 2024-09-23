package routes

import (
	"go-inventory/controllers"
	"go-inventory/internal/config"

	"github.com/go-chi/chi"
)

func Routes(v1Router *chi.Mux, apiCfg *config.ApiConfig) {
	userController := &controllers.UserController{Config: apiCfg} // Ensure this matches the struct definition

	v1Router.Get("/ready", controllers.HandlerReadiness)
	v1Router.Get("/err", controllers.HandlerErr)

	v1Router.Post("/users", userController.HandlerCreateUser)
	v1Router.Get("/users", userController.MiddlewareAuth(userController.HandlerGetUser))

	v1Router.Post("/category", userController.MiddlewareAuth(userController.HandlerCreateCategory))
	v1Router.Delete("/category/{categoryID}", userController.MiddlewareAuth(userController.HandlerDeleteCategory))
	v1Router.Post("/product", userController.MiddlewareAuth(userController.HandlerCreateProduct))

}
