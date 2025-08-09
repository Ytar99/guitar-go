package services

import (
	"guitar-go/internal/models"
	"guitar-go/internal/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
