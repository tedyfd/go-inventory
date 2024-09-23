package main

import (
	"database/sql"
	"fmt"
	"go-inventory/internal/config"
	"go-inventory/routes"
	"log"
	"net/http"
	"os"

	"go-inventory/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"

	//for postgresql driver.
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the Environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the Environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	db := database.New(conn)
	apiCfg := config.ApiConfig{
		DB: db,
	}

	router := chi.NewRouter()
	chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	routes.Routes(v1Router, &apiCfg)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server Run on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port:", portString)
}
