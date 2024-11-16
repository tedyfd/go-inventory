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

	v1Router.Post("/register", userController.HandlerRegisterUser)
	v1Router.Post("/login", userController.HandlerLogin)
	v1Router.Post("/logout", userController.HandlerLogout)

	v1Router.Post("/users", userController.HandlerCreateUser)
	v1Router.Get("/users", userController.JwtAuth(userController.HandlerGetUser))

	v1Router.Post("/category", userController.JwtAuth(userController.HandlerCreateCategory))
	v1Router.Delete("/category/{categoryID}", userController.JwtAuth(userController.HandlerDeleteCategory))
	v1Router.Post("/product", userController.JwtAuth(userController.HandlerCreateProduct))

}
