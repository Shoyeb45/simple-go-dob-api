package service

import "github.com/Shoyeb45/simple-go-dob-api/internal/repository"

type UserService struct {
	repo *repository.UserRepository
}

func (s *UserService) GetUser(id string)  error {

}

