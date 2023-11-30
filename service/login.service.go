package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/depublic/depublic/internal/config"
	"github.com/depublic/depublic/internal/repository"
	"github.com/depublic/depublic/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

// LoginService is a struct that holds the handlers for login-related requests.
type LoginService struct {
	config *config.Config
	repo   repository.UserRepository
}

// NewLoginService returns a new LoginService instance.
func NewLoginService(config *config.Config, repo repository.UserRepository) *LoginService {
	return &LoginService{
		config: config,
		repo:   repo,
	}
}

// LoginHandler handles the `POST /auth/login` request.
func (s *LoginService) LoginHandler(w http.ResponseWriter, r *http.Request) {
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
	user, err := s.repo.GetUserByEmail(context.Background(), loginRequest.Email)
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
		Subject:   user.ID,
		Issuer:    s.config.JWTIssuer,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	})
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
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
