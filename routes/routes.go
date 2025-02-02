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

	v1Router.Get("/users", userController.JwtAuth(userController.HandlerGetUser))

	v1Router.Post("/category", userController.JwtAuth(userController.HandlerCreateCategory))
	v1Router.Get("/category", userController.JwtAuth(userController.HandlerGetCategory))
	v1Router.Get("/category/{categoryName}", userController.JwtAuth(userController.HandlerGetCategoryByName))
	v1Router.Put("/category/{categoryID}", userController.JwtAuth(userController.HandlerUpdateCategory))
	v1Router.Delete("/category/{categoryID}", userController.JwtAuth(userController.HandlerDeleteCategory))

	v1Router.Post("/product", userController.JwtAuth(userController.HandlerCreateProduct))
	v1Router.Get("/product", userController.JwtAuth(userController.HandlerGetProduct))
	v1Router.Delete("/product/{productID}", userController.JwtAuth(userController.HandlerDeleteProduct))

	v1Router.Post("/customer", userController.JwtAuth(userController.HandlerCreateCustomer))
	v1Router.Get("/customer", userController.JwtAuth(userController.HandlerGetCustomer))
	v1Router.Get("/customer/{customerName}", userController.JwtAuth(userController.HandlerGetCustomerByName))
	v1Router.Put("/customer/{customerID}", userController.JwtAuth(userController.HandlerUpdateCustomer))
	v1Router.Delete("/customer/{customerID}", userController.JwtAuth(userController.HandlerDeleteCustomer))

	v1Router.Post("/seller", userController.JwtAuth(userController.HandlerCreateSeller))
	v1Router.Get("/seller", userController.JwtAuth(userController.HandlerGetSeller))
	v1Router.Get("/seller/{sellerName}", userController.JwtAuth(userController.HandlerGetSellerByName))
	v1Router.Put("/seller/{sellerID}", userController.JwtAuth(userController.HandlerUpdateSeller))
	v1Router.Delete("/seller/{sellerID}", userController.JwtAuth(userController.HandlerDeleteSeller))

	v1Router.Post("/order", userController.JwtAuth(userController.HandlerCreateOrder))
}
