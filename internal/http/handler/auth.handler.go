package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/depublic/depublic/internal/config"
	"github.com/depublic/depublic/internal/repository"
	"github.com/depublic/depublic/internal/util"
	"github.com/dgrijalva/jwt-go"
)

// AuthHandler is a struct that holds the handlers for authentication-related requests.
type AuthHandler struct {
	config *config.Config
	repo   repository.Repository
}

// NewAuthHandler returns a new AuthHandler instance.
func NewAuthHandler(config *config.Config, repo repository.Repository) *AuthHandler {
	return &AuthHandler{
		config: config,
		repo:   repo,
	}
}

// LoginHandler handles the `POST /auth/login` request.
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request body
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		util.Error(w, http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}

	// Check if the user exists
	user, err := h.repo.GetUserByEmail(loginRequest.Email)
	if err != nil {
		if err == repository.ErrNotFound {
			util.Error(w, http.StatusBadRequest, "User not found")
		} else {
			util.Error(w, http.StatusInternalServerError, "Failed to get user: %v", err)
		}
		return
	}

	// Check if the password is correct
	if !util.CheckPassword(loginRequest.Password, user.Password) {
		util.Error(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.Claims{
		Subject: user.ID,
		Issuer:  h.config.JWTIssuer,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	})
	tokenString, err := token.SignedString([]byte(h.config.JWTSecret))
	if err != nil {
		util.Error(w, http.StatusInternalServerError, "Failed to generate JWT token: %v", err)
		return
	}

	// Return the JWT token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	})
}

// RegisterHandler handles the `POST /auth/register` request.
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request body
	var registerRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		util.Error(w, http.StatusBadRequest, "Invalid request body: %v", err)
		return
	}

	// Check if the email is already in use
	if _, err := h.repo.GetUserByEmail(registerRequest.Email); err == nil {
		util.Error(w, http.StatusBadRequest, "Email already in use")
		return
	}

	// Hash the password
	hashedPassword, err := util.HashPassword(registerRequest.Password)
	if err != nil {
		util.Error(w, http.StatusInternalServerError, "Failed to hash password: %v", err)
	
