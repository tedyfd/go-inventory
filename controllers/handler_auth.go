package controllers

import (
	"encoding/json"
	"fmt"
	"go-inventory/internal/database"
	"go-inventory/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserController) HandlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(params.Password), 14)

	user, err := uc.Config.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username:  params.Username,
		Password:  string(password),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create User: %v", err))
		return
	}

	respondWithJSON(w, 201, "Register Success", models.DatabaseUserToUser(user))
}

func (uc *UserController) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	username, err := uc.Config.DB.GetUserByUsername(r.Context(), string(params.Username))
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Username no found: %v", err))
		return
	}

	passwordCheck := bcrypt.CompareHashAndPassword([]byte(username.Password), []byte(params.Password))

	if passwordCheck != nil {
		respondWithError(w, 400, fmt.Sprintf("Incorrect Password: %v", err))
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    username.ID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		fmt.Println("Error signing the token:", err)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)

	respondWithJSON(w, 200, "Login success", tokenString)
}

func (uc *UserController) HandlerLogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	respondWithJSON(w, 200, "Success", struct{}{})
}
