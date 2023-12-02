package service

import (
	"depublic/common"
	"depublic/entity"
)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) GenerateToken(claims *entity.JWTClaims) (string, error) {
	token, err := common.GenerateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *TokenService) ValidateToken(tokenString string) (*entity.JWTClaims, error) {
	claims, err := common.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
