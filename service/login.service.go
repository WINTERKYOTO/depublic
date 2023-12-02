package service

import (
	"depublic/entity"
	"depublic/repository"
	"errors"
)

type LoginService struct {
	userRepo repository.UserRepository
}

func NewLoginService(userRepo repository.UserRepository) *LoginService {
	return &LoginService{userRepo}
}

func (s *LoginService) Login(username string, password string) (*entity.User, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
