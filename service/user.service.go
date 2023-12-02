package service

import (
	"depublic/entity"
	"depublic/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) GetAll() ([]*entity.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) GetByID(userID string) (*entity.User, error) {
	return s.userRepo.GetByID(userID)
}

func (s *UserService) Create(user *entity.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) Update(user *entity.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) Delete(userID string) error {
	return s.userRepo.Delete(userID)
}
