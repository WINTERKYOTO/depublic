package service

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// TokenService is a struct that holds the handlers for token-related requests.
type TokenService struct {
	config *config.Config
}

// NewTokenService returns a new TokenService instance.
func NewTokenService(config *config.Config) *TokenService {
	return &TokenService{
		config: config,
	}
}

// ParseToken parses a JWT token.
func (s *TokenService) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTSecret), nil
	})
	if err != nil {
		
