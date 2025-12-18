package service

import "github.com/Shoyeb45/simple-go-dob-api/internal/repository"

type UserService struct {
	repo *repository.UserRepository
}

func (s *UserService) GetUser(id int64)  error {
	return nil;
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo};
}