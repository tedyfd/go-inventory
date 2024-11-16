package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-inventory/internal/auth"
	"go-inventory/internal/database"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtSecret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatalf("Environment variable 'JWT_SECRET' not found")
	} else {
		jwtSecret = []byte(secret)
	}
}

func (uc *UserController) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %w", err))
			return
		}

		user, err := uc.Config.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %w", err))
			return
		}

		handler(w, r, user)
	}
}

func (uc *UserController) JwtAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie("jwt")
		if cookie == nil {
			respondWithError(w, 401, "Unauthorized")
			return
		}
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			fmt.Sprintf("error %v", err)
			respondWithError(w, 401, "Unauthorized")
			return
		}

		issuer, _ := uuid.Parse(claims.Issuer)
		user, _ := uc.Config.DB.GetUserById(r.Context(), issuer)
		handler(w, r, user)
	}
}
