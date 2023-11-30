package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SecretKey is the secret key used to sign and verify JWTs.
var SecretKey = "my-secret-key"

// GenerateToken generates a new JWT with the specified claims.
func GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return tokenString, nil
}

// ParseToken parses a JWT and returns the claims.
func ParseToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return jwt.Claims{}, errors.New("invalid token signature")
		}
		return jwt.Claims{}, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return jwt.Claims{}, errors.New("invalid token")
	}
	return token.Claims, nil
}

// IsTokenExpired checks if a JWT is expired.
func IsTokenExpired(claims jwt.Claims) bool {
	return time.Now().UTC().After(claims.ExpiresAt)
}
